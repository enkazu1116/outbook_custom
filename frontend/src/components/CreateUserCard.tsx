export const CreateUserCard = () => {
    return (
        <div className="create-user-card w-full bg-slate-900/80 border border-slate-700 rounded-2xl px-8 py-6 shadow-[0_0_18px_rgba(34,211,238,0.25)]">
            <h2 className="text-2xl font-bold mb-4 text-cyan-300 drop-shadow-[0_0_10px_rgba(34,211,238,0.6)]">ユーザー作成</h2>
            <form>
                <div className="form-group mb-4">
                    <label htmlFor="name" className="block mb-1 text-slate-200">ユーザー名</label>
                    <input type="text" id="name" name="name" className="w-full px-4 py-2 rounded-lg border border-slate-700 bg-slate-800 text-slate-200 focus:outline-none focus:ring-2 focus:ring-cyan-400" />
                </div>
                <div className="form-group mb-4">
                    <label htmlFor="email" className="block mb-1 text-slate-200">メールアドレス</label>
                    <input type="email" id="email" name="email" className="w-full px-4 py-2 rounded-lg border border-slate-700 bg-slate-800 text-slate-200 focus:outline-none focus:ring-2 focus:ring-cyan-400" />
                </div>
                <div className="form-group mb-4">
                    <label htmlFor="password" className="block mb-1 text-slate-200">パスワード</label>
                    <input type="password" id="password" name="password" className="w-full px-4 py-2 rounded-lg border border-slate-700 bg-slate-800 text-slate-200 focus:outline-none focus:ring-2 focus:ring-cyan-400" />
                </div>
                <button
                    type="submit"
                    className="w-full mt-2 bg-gradient-to-r from-cyan-400 to-cyan-500 text-black font-bold px-4 py-2 rounded-lg shadow-[0_0_12px_rgba(34,211,238,0.7)] hover:brightness-110 transition"
                >
                    作成
                </button>
            </form>
        </div>
    )
}