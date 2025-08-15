export const NotFoundPage = () => {
  return (
    <div className="flex flex-col items-center justify-center min-h-screen p-4 font-sans">
      <div className="text-center max-w-md">
        <h1 className="text-8xl font-bold text-red-500 mb-6 leading-none">404</h1>
        <h2 className="text-2xl font-semibold text-gray-900 dark:text-gray-100 mb-2">页面未找到</h2>
        <p className="text-gray-600 dark:text-gray-400 mb-8 leading-relaxed">抱歉，您访问的页面不存在或已被移除</p>
        <div className="flex gap-4 flex-wrap justify-center">
          <button
            onClick={() => (window.location.href = "/")}
            className="px-6 py-3 bg-blue-600 hover:bg-blue-700 text-white rounded-lg transition-colors cursor-pointer"
          >
            返回首页
          </button>
          <button
            onClick={() => window.history.back()}
            className="px-6 py-3 bg-transparent hover:bg-gray-100 dark:hover:bg-gray-800 text-gray-700 dark:text-gray-300 border border-gray-300 dark:border-gray-600 rounded-lg transition-colors cursor-pointer"
          >
            返回上页
          </button>
        </div>
      </div>
    </div>
  )
}
