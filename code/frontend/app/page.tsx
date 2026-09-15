import { ShowPersistedGreeting } from "../components/ShowPersistedGreeting";
import { getGreeting } from "../lib/show-persisted-greeting";

export default async function Home() {
  const { greeting } = await getGreeting();

  return (
    <main>
      <ShowPersistedGreeting initialGreeting={greeting} />
    </main>
  );
}
