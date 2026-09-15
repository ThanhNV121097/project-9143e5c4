import { EditPersistedGreeting } from "../components/EditPersistedGreeting";
import { getGreeting } from "../lib/edit-persisted-greeting";

export default async function Home() {
  const { greeting } = await getGreeting();

  return (
    <main>
      <EditPersistedGreeting initialGreeting={greeting} />
    </main>
  );
}
