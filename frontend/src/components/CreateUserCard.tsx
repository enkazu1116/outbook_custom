import { useState, type ChangeEvent, type FormEvent } from "react";

type CreateUserForm = {
    name: string;
    email: string;
    password: string;
};

type CreateUserCardProps = {
    onCreate: (input: CreateUserForm) => Promise<void>;
};

export const CreateUserCard = ({ onCreate }: CreateUserCardProps) => {

    // フォームの状態管理
    const [form, setForm] = useState<CreateUserForm>({
        name: '',
        email: '',
        password: '',
    });
    const [isSubmitting, setIsSubmitting] = useState(false);
    const [errorMessage, setErrorMessage] = useState<string | null>(null);

    //　変更を反映
    const handleChange = (
        e: ChangeEvent<HTMLInputElement>
    ) => {
        const { name, value } = e.target;
        setForm((prev) => ({
            ...prev, 
            [name]: value,
        }));
    };

    // フォーム送信
    const handleSubmit = async (
        e: FormEvent<HTMLFormElement>
    ) => {
        e.preventDefault();
        setErrorMessage(null);
        // バリデーション
        if (!form.name || !form.email || !form.password) {
            alert('ユーザー名、メールアドレス、パスワードを入力してください');
            return;
        }

        // backendのドメインルールに合わせて事前チェック（8文字以上・半角英数字）
        if (form.password.length < 8) {
            setErrorMessage("パスワードは8文字以上で入力してください");
            return;
        }
        if (!/^[a-zA-Z0-9]+$/.test(form.password)) {
            setErrorMessage("パスワードは半角英数字のみで入力してください");
            return;
        }

        setIsSubmitting(true);
        try {
            await onCreate(form);
            setForm({ name: "", email: "", password: "" });
        } catch (e) {
            const message =
                e instanceof Error ? e.message : "ユーザー作成に失敗しました";
            setErrorMessage(message);
        } finally {
            setIsSubmitting(false);
        }
    };
    return (
        <div className="create-user-card w-full bg-slate-900/80 border border-slate-700 rounded-2xl px-8 py-6 shadow-[0_0_18px_rgba(34,211,238,0.25)]">
            <h2 className="text-2xl font-bold mb-4 text-cyan-300 drop-shadow-[0_0_10px_rgba(34,211,238,0.6)]">
                ユーザー作成
            </h2>
            {errorMessage ? (
                <div className="mb-4 rounded-lg border border-rose-500/60 bg-rose-950/40 px-4 py-2 text-rose-200">
                    {errorMessage}
                </div>
            ) : null}
            {/* フォーム */}
            <form onSubmit={handleSubmit}>
                <div className="form-group mb-4">
                    <label htmlFor="name" className="block mb-1 text-slate-200">
                        ユーザー名
                    </label>
                    <input 
                        type="text" 
                        id="name" 
                        name="name" 
                        className="w-full px-4 py-2 rounded-lg border border-slate-700 bg-slate-800 text-slate-200 focus:outline-none focus:ring-2 focus:ring-cyan-400" 
                        value={form.name}
                        onChange={handleChange}
                    />
                </div>
                <div className="form-group mb-4">
                    <label htmlFor="email" className="block mb-1 text-slate-200">
                        メールアドレス
                    </label>
                    <input 
                        type="email" 
                        id="email" 
                        name="email" 
                        className="w-full px-4 py-2 rounded-lg border border-slate-700 bg-slate-800 text-slate-200 focus:outline-none focus:ring-2 focus:ring-cyan-400" 
                        value={form.email}
                        onChange={handleChange}
                    />
                </div>
                <div className="form-group mb-4">
                    <label htmlFor="password" className="block mb-1 text-slate-200">
                        パスワード
                    </label>
                    <input 
                        type="password" 
                        id="password" 
                        name="password" 
                        className="w-full px-4 py-2 rounded-lg border border-slate-700 bg-slate-800 text-slate-200 focus:outline-none focus:ring-2 focus:ring-cyan-400" 
                        value={form.password}
                        onChange={handleChange}
                    />
                </div>
                <button
                    type="submit"
                    disabled={isSubmitting}
                    className="w-full mt-2 bg-gradient-to-r from-cyan-400 to-cyan-500 text-black font-bold px-4 py-2 rounded-lg shadow-[0_0_12px_rgba(34,211,238,0.7)] hover:brightness-110 transition disabled:opacity-60 disabled:cursor-not-allowed"
                >
                    {isSubmitting ? "作成中..." : "作成"}
                </button>
            </form>
        </div>
    )
}