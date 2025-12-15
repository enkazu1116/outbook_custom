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
  const handleCreateUser = (input: { name: string; email: string; password: string }) => {
    // TODO: API 呼び出し
    setUsers((prev) => [
      ...prev, 
      { id: Date.now().toString(), 
        ...input, 
        createdAt: new Date().toISOString(), 
        updatedAt: new Date().toISOString() 
      }]);
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
            <UsersTable />
            <CreateUserCard onCreate={handleCreateUser} />
          </div>

        </main>
      </div>
    </>
  )
}

export default App
