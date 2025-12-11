export const BulkOperations = () => {
    return (
        <div className="bulk-operations w-full mt-10 px-8 flex gap-4 bg-slate-900/60 border border-slate-700 rounded-xl py-4 shadow-[0_0_15px_rgba(34,211,238,0.2)]">
            <button
                type="button"
                className="px-6 py-3 rounded-xl bg-cyan-500 !bg-cyan-500 text-white font-semibold shadow-[0_0_12px_rgba(34,211,238,0.6)] hover:bg-cyan-400 hover:shadow-[0_0_18px_rgba(34,211,238,0.9)] active:bg-cyan-600 active:shadow-[0_0_8px_rgba(34,211,238,0.5)] transition-all duration-150"
            >
                Bulk Operations Toolbar
            </button>
            <button
                type="button"
                className="px-6 py-3 rounded-xl bg-cyan-500 !bg-cyan-500 text-white font-semibold shadow-[0_0_12px_rgba(34,211,238,0.6)] hover:bg-cyan-400 hover:shadow-[0_0_18px_rgba(34,211,238,0.9)] active:bg-cyan-600 active:shadow-[0_0_8px_rgba(34,211,238,0.5)] transition-all duration-150"
            >
                Bulk Status Update
            </button>
            <button
                type="button"
                className="px-6 py-3 rounded-xl bg-cyan-500 !bg-cyan-500 text-white font-semibold shadow-[0_0_12px_rgba(34,211,238,0.6)] hover:bg-cyan-400 hover:shadow-[0_0_18px_rgba(34,211,238,0.9)] active:bg-cyan-600 active:shadow-[0_0_8px_rgba(34,211,238,0.5)] transition-all duration-150"
            >
                Disable Loop
            </button>
        </div>
    )
}