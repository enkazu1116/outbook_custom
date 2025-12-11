import { Header } from './components/Header'
import { SearchPanel } from './components/SearchPanel'
import { BulkOperations } from './components/BulkOperations'
import { UsersTable } from './components/UsersTable'
import { CreateUserCard } from './components/CreateUserCard'

function App() {

  return (
    <>
      <Header />
      <div className="min-h-screen bg-slate-950 text-white pt-5 w-screen overflow-x-hidden">
        <div className="px-8 mt-2 w-screen">
          <SearchPanel />
        </div>
        <main className="w-screen px-8 mt-4 space-y-6">

          {/* BulkOperations alone */}
          <BulkOperations />

          {/* UsersTable + CreateUserCard side-by-side */}
          <div className="grid grid-cols-1 lg:grid-cols-[2fr_1fr] gap-6">
            <UsersTable />
            <CreateUserCard />
          </div>

        </main>
      </div>
    </>
  )
}

export default App
