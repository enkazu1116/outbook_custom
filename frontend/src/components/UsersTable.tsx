export const UsersTable = () => {
    return (
        <div className="users-table w-full mt-8 px-8">
            <table className="w-full border-collapse bg-slate-900/60 rounded-xl overflow-hidden shadow-[0_0_18px_rgba(34,211,238,0.15)]">
                <thead className="bg-slate-800/80 border-b border-slate-700">
                    <tr>
                        <th className="text-left px-4 py-3 text-slate-200">ユーザー名</th>
                        <th className="text-left px-4 py-3 text-slate-200">メールアドレス</th>
                        <th className="text-left px-4 py-3 text-slate-200">ステータス</th>
                        <th className="text-center px-4 py-3 text-slate-200">編集</th>
                        <th className="text-center px-4 py-3 text-slate-200">削除</th>
                    </tr>
                </thead>
                <tbody>
                    <tr className="border-b border-slate-700/40 hover:bg-slate-800/50 transition">
                        <td className="px-4 py-3 text-slate-200">ユーザー名</td>
                        <td className="px-4 py-3 text-slate-200">メールアドレス</td>
                        <td className="px-4 py-3 text-slate-200">アクティブ</td>
                        <td className="px-4 py-3 text-center">
                            <button className="px-3 py-1 rounded bg-emerald-500 !bg-emerald-500 text-white font-bold shadow-[0_0_18px_rgba(16,185,129,1)] hover:shadow-[0_0_25px_rgba(16,185,129,1)] hover:brightness-110 transition">
                                編集
                            </button>
                        </td>
                        <td className="px-4 py-3 text-center">
                            <button className="px-3 py-1 rounded bg-rose-600 !bg-rose-600 text-white font-bold shadow-[0_0_18px_rgba(244,63,94,1)] hover:shadow-[0_0_25px_rgba(244,63,94,1)] hover:brightness-110 transition">
                                削除
                            </button>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    );
}