import { useState, useEffect } from 'react'

// Types
interface Plugin {
  id: string
  name: string
  version: string
  status: string
}

interface Theme {
  id: string
  name: string
  version: string
  is_active: boolean
  type: string
}

interface Stats {
  totalPlugins: number
  runningPlugins: number
  totalThemes: number
  activeTheme: string
}

function App() {
  const [currentPage, setCurrentPage] = useState('dashboard')
  const [plugins, setPlugins] = useState<Plugin[]>([])
  const [themes, setThemes] = useState<Theme[]>([])
  const [stats, setStats] = useState<Stats>({
    totalPlugins: 0,
    runningPlugins: 0,
    totalThemes: 0,
    activeTheme: 'moon'
  })

  // Fetch data from APIs
  useEffect(() => {
    fetchPlugins()
    fetchThemes()
  }, [])

  const fetchPlugins = async () => {
    try {
      const response = await fetch('/api/plugin/list?page_no=1&page_size=100')
      if (response.ok) {
        const data = await response.json()
        setPlugins(data.data?.list || [])
        setStats(prev => ({
          ...prev,
          totalPlugins: data.data?.total || 0,
          runningPlugins: data.data?.list?.filter((p: Plugin) => p.status === 'running').length || 0
        }))
      }
    } catch (error) {
      console.error('Failed to fetch plugins:', error)
    }
  }

  const fetchThemes = async () => {
    try {
      const response = await fetch('/api/theme/list?page_no=1&page_size=100')
      if (response.ok) {
        const data = await response.json()
        setThemes(data.data?.list || [])
        const activeTheme = data.data?.list?.find((t: Theme) => t.is_active)
        setStats(prev => ({
          ...prev,
          totalThemes: data.data?.total || 0,
          activeTheme: activeTheme?.name || 'moon'
        }))
      }
    } catch (error) {
      console.error('Failed to fetch themes:', error)
    }
  }

  const switchTheme = async (themeId: string) => {
    try {
      const currentActiveTheme = themes.find(t => t.is_active);
      if (currentActiveTheme?.id === themeId) {
        return;
      }

      const targetTheme = themes.find(t => t.id === themeId);
      const needsRebuild = targetTheme?.id === 'com.jank.themes.moon';

      // 从目标主题配置中获取主题类型
      const themeType = targetTheme?.type;

      const response = await fetch('/api/theme/switch', { 
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ 
          id: themeId,
          theme_type: themeType,
          rebuild: needsRebuild
        })
      });
      
      if (response.ok) {
        const data = await response.json();
        console.log('Theme switch response:', data);
        
        setTimeout(() => {
          window.location.reload();
        }, 500);
      }
    } catch (error) {
      console.error('Failed to switch theme:', error)
    }
  }

  return (
    <div className="min-h-screen bg-gray-900">
      {/* Header */}
      <header className="bg-gray-800 border-b border-gray-700">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            <div className="flex items-center space-x-4">
              <div className="w-8 h-8 bg-blue-600 rounded-lg flex items-center justify-center">
                <span className="text-white font-bold">🌙</span>
              </div>
              <h1 className="text-xl font-semibold text-white">Jank Blog - Moon Theme</h1>
            </div>
            <nav className="flex items-center space-x-4">
              <button
                onClick={() => setCurrentPage('dashboard')}
                className={`px-3 py-2 rounded-md text-sm font-medium ${
                  currentPage === 'dashboard'
                    ? 'bg-blue-600 text-white'
                    : 'text-slate-300 hover:text-white hover:bg-slate-800'
                }`}
              >
                仪表盘
              </button>
              <button
                onClick={() => setCurrentPage('plugins')}
                className={`px-3 py-2 rounded-md text-sm font-medium ${
                  currentPage === 'plugins'
                    ? 'bg-blue-600 text-white'
                    : 'text-slate-300 hover:text-white hover:bg-slate-800'
                }`}
              >
                插件管理
              </button>
              <button
                onClick={() => setCurrentPage('themes')}
                className={`px-3 py-2 rounded-md text-sm font-medium ${
                  currentPage === 'themes'
                    ? 'bg-blue-600 text-white'
                    : 'text-slate-300 hover:text-white hover:bg-slate-800'
                }`}
              >
                主题管理
              </button>
              <a
                href="/console"
                className="px-3 py-2 bg-green-600 hover:bg-green-700 text-white text-sm font-medium rounded-md transition-colors flex items-center space-x-1"
              >
                <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path>
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                </svg>
                <span>管理后台</span>
              </a>
            </nav>
          </div>
        </div>
      </header>

      {/* Main Content */}
      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {currentPage === 'dashboard' && (
          <div className="space-y-6">
            <h2 className="text-2xl font-bold text-white">系统概览</h2>
            
            {/* Stats Cards */}
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
              <div className="card">
                <h3 className="text-sm font-medium text-gray-400">插件总数</h3>
                <p className="text-2xl font-bold text-white">{stats.totalPlugins}</p>
              </div>
              <div className="card">
                <h3 className="text-sm font-medium text-gray-400">运行中插件</h3>
                <p className="text-2xl font-bold text-green-400">{stats.runningPlugins}</p>
              </div>
              <div className="card">
                <h3 className="text-sm font-medium text-gray-400">主题总数</h3>
                <p className="text-2xl font-bold text-white">{stats.totalThemes}</p>
              </div>
              <div className="card">
                <h3 className="text-sm font-medium text-gray-400">当前主题</h3>
                <p className="text-2xl font-bold text-blue-400">{stats.activeTheme}</p>
              </div>
            </div>
          </div>
        )}

        {currentPage === 'plugins' && (
          <div className="space-y-6">
            <h2 className="text-2xl font-bold text-white">插件管理</h2>
            <div className="grid gap-4">
              {plugins.map((plugin) => (
                <div key={plugin.id} className="card">
                  <div className="flex justify-between items-center">
                    <div>
                      <h3 className="text-lg font-medium text-white">{plugin.name}</h3>
                      <p className="text-sm text-gray-400">版本: {plugin.version}</p>
                    </div>
                    <span className={`px-2 py-1 rounded text-xs font-medium ${
                      plugin.status === 'running' 
                        ? 'bg-green-100 text-green-800' 
                        : 'bg-gray-100 text-gray-800'
                    }`}>
                      {plugin.status}
                    </span>
                  </div>
                </div>
              ))}
              {plugins.length === 0 && (
                <div className="card text-center">
                  <p className="text-gray-400">暂无插件数据</p>
                </div>
              )}
            </div>
          </div>
        )}

        {currentPage === 'themes' && (
          <div className="space-y-6">
            <h2 className="text-2xl font-bold text-white">主题管理</h2>
            <div className="grid gap-4">
              {themes.map((theme) => (
                <div key={theme.id} className="card">
                  <div className="flex justify-between items-center">
                    <div>
                      <h3 className="text-lg font-medium text-white">{theme.name}</h3>
                      <p className="text-sm text-gray-400">版本: {theme.version}</p>
                    </div>
                    <div className="flex items-center space-x-2">
                      {theme.is_active && (
                        <span className="px-2 py-1 bg-green-100 text-green-800 rounded text-xs font-medium">
                          当前激活
                        </span>
                      )}
                      {!theme.is_active && (
                        <button
                          onClick={() => switchTheme(theme.id)}
                          className="btn-primary text-sm"
                        >
                          切换
                        </button>
                      )}
                    </div>
                  </div>
                </div>
              ))}
              {themes.length === 0 && (
                <div className="card text-center">
                  <p className="text-gray-400">暂无主题数据</p>
                </div>
              )}
            </div>
          </div>
        )}
      </main>
    </div>
  )
}

export default App
