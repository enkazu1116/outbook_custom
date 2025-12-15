import { Header } from './components/Header'
import { SearchPanel } from './components/SearchPanel'
import { BulkOperations } from './components/BulkOperations'
import { UsersTable } from './components/UsersTable'
import { CreateUserCard } from './components/CreateUserCard'
import { useState } from 'react';

// ユーザー
type User = {
  id: string;
  name: string;
  email: string;
  password: string;
  createdAt: string;
  updatedAt: string;
}

function App() {

  const [users, setUsers] = useState<User[]>([]);
  const handleCreateUser = async (input: { name: string; email: string; password: string }) => {
    const res = await fetch('/users', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        ...input,
        bio: '',
      }),
    });

    if (!res.ok) {
      // backendは {error: "..."} を返す想定
      let message = 'ユーザー作成に失敗しました';
      try {
        const data = (await res.json()) as { error?: string };
        if (data?.error) message = data.error;
      } catch {
        // ignore
      }
      throw new Error(message);
    }

    // 現状 backend が作成したユーザーを返さないため、UIは仮の1件を追加（必要なら後で一覧API実装へ）
    setUsers((prev) => [
      ...prev,
      {
        id: Date.now().toString(),
        ...input,
        createdAt: new Date().toISOString(),
        updatedAt: new Date().toISOString(),
      },
    ]);
  };

  return (
    <>
      <Header />
      <div className="min-h-screen bg-slate-950 text-white pt-5 w-screen overflow-x-hidden">
        <div className="px-8 mt-2 w-screen">
          <SearchPanel />
        </div>
        <main className="w-screen px-8 mt-4 space-y-6">
          <BulkOperations />

          <div className="grid grid-cols-1 lg:grid-cols-[2fr_1fr] gap-6">
            <UsersTable users={users} />
            <CreateUserCard onCreate={handleCreateUser} />
          </div>

        </main>
      </div>
    </>
  )
}

export default App
