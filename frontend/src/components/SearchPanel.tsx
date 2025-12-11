export const SearchPanel = () => {
    return (
        <div className="user-search-panel w-full mt-28 px-8">
            <form
                className="w-full bg-slate-900/80 border border-slate-700 rounded-xl px-6 py-5 flex items-center gap-4 shadow-[0_0_12px_rgba(34,211,238,0.25)]"
            >
                <input
                    type="text"
                    placeholder="ユーザー名を入力"
                    className="flex-1 bg-slate-800 text-slate-200 placeholder-slate-500 px-4 py-2 rounded-lg focus:outline-none focus:ring-2 focus:ring-cyan-400"
                />
                <button
                    type="submit"
                    className="bg-gradient-to-r from-cyan-400 to-cyan-500 text-black font-bold px-5 py-2 rounded-lg shadow-[0_0_10px_rgba(34,211,238,0.7)] hover:brightness-110 transition"
                >
                    検索
                </button>
            </form>
        </div>
    )
}