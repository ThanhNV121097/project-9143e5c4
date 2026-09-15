import { EditPersistedGreeting } from "../components/EditPersistedGreeting";
import { getMockGreeting } from "../lib/mock/edit-persisted-greeting";

export default function Home() {
  const { greeting } = getMockGreeting();

  return (
    <main>
      <EditPersistedGreeting initialGreeting={greeting} />
    </main>
  );
}
