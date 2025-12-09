import { Header } from './components/Header'
import { SearchPanel } from './components/SearchPanel'
import { BulkOperations } from './components/BulkOperations'
import { UsersTable } from './components/UsersTable'
import { CreateUserCard } from './components/CreateUserCard'

function App() {

  return (
    <>
      <Header />
      <SearchPanel />
      <main>
        <section>
          <BulkOperations />
          <UsersTable />
        </section>

        <aside>
          <CreateUserCard />
        </aside>
      </main>
    </>
  )
}

export default App
