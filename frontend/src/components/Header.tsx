export const Header = () => {
    return (
        <header className="outbook-header w-full fixed top-0 left-0 right-0 z-50 bg-gradient-to-r from-slate-900 via-slate-950 to-slate-900 px-10 py-6 flex items-center shadow-lg">
            <h1
                className="text-4xl md:text-5xl lg:text-6xl font-extrabold tracking-wide text-transparent bg-clip-text bg-gradient-to-r from-cyan-300 via-cyan-400 to-cyan-500 drop-shadow-[0_0_25px_rgba(34,211,238,1)]"
            >
                Outbook
            </h1>
        </header>
    );
};