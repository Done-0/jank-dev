// Modern JavaScript for Jank Blog Server Admin Panel
// 使用真实API调用获取数据

class AdminPanel {
    constructor() {
        this.currentSection = 'dashboard';
        this.apiBase = '/api';
        this.themes = []; // 存储主题列表
        this.init();
    }

    init() {
        this.setupNavigation();
        this.loadInitialData();
        this.setupEventListeners();
    }

    // 导航管理
    setupNavigation() {
        const navItems = document.querySelectorAll('.nav-item');
        const sections = document.querySelectorAll('.content-section');

        navItems.forEach(item => {
            item.addEventListener('click', (e) => {
                e.preventDefault();
                const sectionName = item.dataset.section;
                this.switchSection(sectionName);
            });
        });
    }

    switchSection(sectionName) {
        // 更新导航状态
        document.querySelectorAll('.nav-item').forEach(item => {
            item.classList.remove('bg-blue-50', 'dark:bg-blue-900/20', 'text-blue-600', 'dark:text-blue-400');
            item.classList.add('text-gray-700', 'dark:text-gray-300');
        });
        
        const activeItem = document.querySelector(`[data-section="${sectionName}"]`);
        if (activeItem) {
            activeItem.classList.remove('text-gray-700', 'dark:text-gray-300');
            activeItem.classList.add('bg-blue-50', 'dark:bg-blue-900/20', 'text-blue-600', 'dark:text-blue-400');
        }

        // 更新内容区域
        document.querySelectorAll('.content-section').forEach(section => {
            section.classList.add('hidden');
        });
        
        const activeSection = document.getElementById(sectionName);
        if (activeSection) {
            activeSection.classList.remove('hidden');
        }

        this.currentSection = sectionName;

        // 根据当前页面加载数据
        if (sectionName === 'plugins') {
            this.loadPlugins();
        } else if (sectionName === 'themes') {
            this.loadThemes();
        }
    }

    // API调用方法
    async apiCall(endpoint, options = {}) {
        try {
            const response = await fetch(`${this.apiBase}${endpoint}`, {
                headers: {
                    'Content-Type': 'application/json',
                    ...options.headers
                },
                ...options
            });

            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            return await response.json();
        } catch (error) {
            console.error('API call failed:', error);
            this.showToast('API调用失败', error.message, 'error');
            throw error;
        }
    }

    // 加载初始数据
    async loadInitialData() {
        await Promise.all([
            this.loadDashboardStats(),
            this.loadActiveTheme()
        ]);
    }

    // 加载仪表盘统计数据
    async loadDashboardStats() {
        try {
            const [pluginsData, themesData] = await Promise.all([
                this.apiCall('/plugin/list?page_no=1&page_size=100'),
                this.apiCall('/theme/list?page_no=1&page_size=100')
            ]);

            // 更新插件统计
            const totalPlugins = pluginsData.data?.total || 0;
            const runningPlugins = pluginsData.data?.list?.filter(p => p.status === 'running').length || 0;
            
            document.getElementById('total-plugins').textContent = totalPlugins;
            document.getElementById('running-plugins').textContent = runningPlugins;
            document.getElementById('plugin-badge').textContent = totalPlugins;

            // 更新主题统计
            const totalThemes = themesData.data?.total || 0;
            document.getElementById('total-themes').textContent = totalThemes;
            document.getElementById('theme-badge').textContent = totalThemes;

        } catch (error) {
            console.error('Failed to load dashboard stats:', error);
            // 设置默认值
            document.getElementById('total-plugins').textContent = '0';
            document.getElementById('running-plugins').textContent = '0';
            document.getElementById('total-themes').textContent = '0';
            document.getElementById('plugin-badge').textContent = '0';
            document.getElementById('theme-badge').textContent = '0';
        }
    }

    // 加载当前激活主题
    async loadActiveTheme() {
        try {
            const response = await this.apiCall('/theme/get');
            const activeTheme = response.data?.theme;
            
            if (activeTheme) {
                document.getElementById('active-theme-name').textContent = activeTheme.name || 'Unknown';
            } else {
                document.getElementById('active-theme-name').textContent = 'None';
            }
        } catch (error) {
            console.error('Failed to load active theme:', error);
            document.getElementById('active-theme-name').textContent = 'Error';
        }
    }

    // 加载插件列表
    async loadPlugins() {
        const loadingEl = document.getElementById('plugins-loading');
        const contentEl = document.getElementById('plugins-content');
        const emptyEl = document.getElementById('plugins-empty');

        // 显示加载状态
        loadingEl.classList.remove('hidden');
        contentEl.classList.add('hidden');
        emptyEl.classList.add('hidden');

        try {
            const response = await this.apiCall('/plugin/list?page_no=1&page_size=50');
            const plugins = response.data?.list || [];

            if (plugins.length === 0) {
                loadingEl.classList.add('hidden');
                emptyEl.classList.remove('hidden');
                return;
            }

            // 渲染插件卡片
            contentEl.innerHTML = plugins.map(plugin => this.renderPluginCard(plugin)).join('');
            
            loadingEl.classList.add('hidden');
            contentEl.classList.remove('hidden');

        } catch (error) {
            console.error('Failed to load plugins:', error);
            loadingEl.classList.add('hidden');
            emptyEl.classList.remove('hidden');
        }
    }

    // 渲染插件卡片
    renderPluginCard(plugin) {
        const statusClass = this.getStatusClass(plugin.status);
        const statusText = this.getStatusText(plugin.status);

        return `
        <div class="bg-white dark:bg-gray-800 rounded-lg border border-gray-200 dark:border-gray-700 p-6 shadow-sm hover:shadow-md transition-all h-full flex flex-col">
            <div class="flex items-center justify-between mb-4">
                <div class="w-12 h-12 bg-gradient-to-br from-blue-400 to-indigo-500 rounded-lg flex items-center justify-center text-white text-xl font-bold">
                    ${(plugin.name || plugin.id).charAt(0).toUpperCase()}
                </div>
                <span class="px-2 py-1 text-xs font-medium rounded-full ${statusClass}">${statusText}</span>
            </div>
            
            <div class="mb-4 flex-1">
                <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-1">${plugin.name || plugin.id}</h3>
                <p class="text-sm text-gray-600 dark:text-gray-400 mb-2">版本: v${plugin.version || '1.0.0'}</p>
                <p class="text-sm text-gray-600 dark:text-gray-400 mb-2">类型: ${plugin.type || 'unknown'}</p>
                <p class="text-sm text-gray-600 dark:text-gray-400 mb-2">作者: ${plugin.author || 'Unknown'}</p>
                <p class="text-sm text-gray-500 dark:text-gray-400">${plugin.description || '暂无描述'}</p>
            </div>
        </div>
    `;
    }

    // 加载主题列表
    async loadThemes() {
        const loadingEl = document.getElementById('themes-loading');
        const contentEl = document.getElementById('themes-content');
        const emptyEl = document.getElementById('themes-empty');

        // 显示加载状态
        loadingEl.style.display = 'block';
        contentEl.style.display = 'none';
        emptyEl.style.display = 'none';

        try {
            const response = await this.apiCall('/theme/list?page_no=1&page_size=50');
            const themes = response.data?.list || [];
            
            // 存储主题列表到实例变量
            this.themes = themes;

            if (themes.length === 0) {
                loadingEl.style.display = 'none';
                emptyEl.style.display = 'block';
                return;
            }

            // 渲染主题卡片
            contentEl.innerHTML = themes.map(theme => this.renderThemeCard(theme)).join('');
            
            // 添加主题切换事件监听器
            this.setupThemeClickHandlers();
            
            loadingEl.style.display = 'none';
            contentEl.style.display = 'grid';

        } catch (error) {
            console.error('Failed to load themes:', error);
            loadingEl.style.display = 'none';
            emptyEl.style.display = 'block';
        }
    }

    // 渲染主题卡片
    renderThemeCard(theme) {
        const isActive = theme.is_active;
        const activeClass = isActive ? 'border-blue-500 bg-blue-50 dark:bg-blue-900/20' : 'border-gray-200 dark:border-gray-700';
        const statusClass = isActive ? 'bg-green-100 text-green-800 dark:bg-green-900/20 dark:text-green-400' : 'bg-gray-100 text-gray-800 dark:bg-gray-900/20 dark:text-gray-400';
        const statusText = isActive ? '当前主题' : '可切换';
        const buttonClass = isActive ? 'bg-gray-400 cursor-not-allowed' : 'bg-blue-600 hover:bg-blue-700 cursor-pointer';
        const buttonText = isActive ? '当前使用' : '切换主题';
        const buttonDisabled = isActive ? 'disabled' : '';

        return `
        <div class="bg-white dark:bg-gray-800 rounded-lg border ${activeClass} p-6 shadow-sm hover:shadow-md transition-all theme-card h-full flex flex-col" data-theme-id="${theme.id}">
            <div class="flex items-center justify-between mb-4">
                <div class="w-12 h-12 bg-gradient-to-br from-purple-400 to-blue-500 rounded-lg flex items-center justify-center text-white text-xl font-bold">
                    ${(theme.name || theme.id).charAt(0).toUpperCase()}
                </div>
                <span class="px-2 py-1 text-xs font-medium rounded-full ${statusClass}">${statusText}</span>
            </div>
            
            <div class="mb-4 flex-1">
                <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-1">${theme.name || theme.id}</h3>
                <p class="text-sm text-gray-600 dark:text-gray-400 mb-2">作者: ${theme.author || 'Unknown'}</p>
                <p class="text-sm text-gray-600 dark:text-gray-400 mb-2">版本: ${theme.version || 'N/A'}</p>
                <p class="text-sm text-gray-500 dark:text-gray-400">${theme.description || '暂无描述'}</p>
            </div>
            
            <button 
                class="w-full px-4 py-2 text-white text-sm font-medium rounded-lg transition-colors ${buttonClass} theme-switch-btn mt-auto"
                data-theme-id="${theme.id}"
                ${buttonDisabled}
            >
                ${buttonText}
            </button>
        </div>
    `;
    }

    // 设置主题点击处理器
    setupThemeClickHandlers() {
        const themeSwitchBtns = document.querySelectorAll('.theme-switch-btn');
        themeSwitchBtns.forEach(btn => {
            btn.addEventListener('click', (e) => {
                e.stopPropagation(); // 防止事件冒泡
                const themeId = btn.dataset.themeId;
                if (!btn.disabled) {
                    this.switchTheme(themeId);
                }
            });
        });
    }

    // 切换主题
    async switchTheme(themeId) {
        try {
            // 获取当前激活的主题
            const currentThemeResponse = await this.apiCall('/theme/get');
            const currentThemeId = currentThemeResponse.data?.theme?.id;
            
            // 如果要切换的主题就是当前主题，不需要操作
            if (currentThemeId === themeId) {
                this.showToast('主题未更改', '当前已是该主题', 'info');
                return;
            }
            
            // 从目标主题配置中获取主题类型
            const targetTheme = this.themes.find(theme => theme.id === themeId);
            const themeType = targetTheme?.type || 'frontend';
            
            this.showToast('正在切换主题...', '请稍候', 'info');
            
            console.log('Switching to theme:', themeId, 'type:', themeType);
            
            // 使用AJAX调用，后端返回成功后刷新页面
            const response = await this.apiCall('/theme/switch', {
                method: 'POST',
                body: JSON.stringify({ 
                    id: themeId,
                    theme_type: themeType
                })
            });
            
            console.log('Theme switch response:', response);
            
            // 主题切换成功，直接刷新页面应用新主题
            this.showToast('主题切换成功', '页面即将刷新', 'success');
            setTimeout(() => {
                window.location.reload();
            }, 500);
        } catch (error) {
            console.error('Failed to switch theme:', error);
            this.showToast('主题切换失败', error.message, 'error');
        }
    }

    // 获取状态样式类
    getStatusClass(status) {
        const statusMap = {
            'running': 'bg-green-100 text-green-800 dark:bg-green-900/20 dark:text-green-400',
            'loaded': 'bg-yellow-100 text-yellow-800 dark:bg-yellow-900/20 dark:text-yellow-400',
            'stopped': 'bg-red-100 text-red-800 dark:bg-red-900/20 dark:text-red-400',
            'error': 'bg-red-100 text-red-800 dark:bg-red-900/20 dark:text-red-400',
            'ready': 'bg-blue-100 text-blue-800 dark:bg-blue-900/20 dark:text-blue-400'
        };
        return statusMap[status] || 'bg-gray-100 text-gray-800 dark:bg-gray-900/20 dark:text-gray-400';
    }

    // 获取状态文本
    getStatusText(status) {
        const statusMap = {
            'running': '运行中',
            'loaded': '已加载',
            'stopped': '已停止',
            'error': '错误',
            'ready': '就绪'
        };
        return statusMap[status] || status || '未知';
    }

    // Toast通知系统
    showToast(title, message, type = 'info') {
        const toastContainer = document.getElementById('toast-container');
        const toastId = 'toast-' + Date.now();
        
        const iconMap = {
            'success': '✅',
            'error': '❌',
            'warning': '⚠️',
            'info': 'ℹ️'
        };

        const toast = document.createElement('div');
        toast.id = toastId;
        toast.className = `toast ${type}`;
        toast.innerHTML = `
            <div class="toast-icon">${iconMap[type]}</div>
            <div class="toast-content">
                <div class="toast-title">${title}</div>
                <div class="toast-message">${message}</div>
            </div>
            <button class="toast-close" onclick="this.parentElement.remove()">×</button>
        `;

        toastContainer.appendChild(toast);

        // 自动移除Toast
        setTimeout(() => {
            const toastEl = document.getElementById(toastId);
            if (toastEl) {
                toastEl.remove();
            }
        }, 5000);
    }

    // 设置事件监听器
    setupEventListeners() {
        // 全局刷新按钮
        window.refreshPlugins = () => this.loadPlugins();
        window.refreshThemes = () => this.loadThemes();
        window.switchToSection = (section) => this.switchSection(section);
    }
}

// 初始化应用
document.addEventListener('DOMContentLoaded', () => {
    new AdminPanel();
});
