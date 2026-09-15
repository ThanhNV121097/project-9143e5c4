import styles from "./ShowPersistedGreeting.module.css";

type ShowPersistedGreetingProps = {
  initialGreeting: string;
};

export function ShowPersistedGreeting({ initialGreeting }: ShowPersistedGreetingProps) {
  return (
    <section className={styles.section} aria-labelledby="greeting-heading">
      <h1 id="greeting-heading" className={styles.heading}>
        {initialGreeting}
      </h1>
      <form className={styles.form} autoComplete="off">
        <label className={styles.label} htmlFor="greeting-input">
          Greeting
        </label>
        <input
          className={styles.input}
          id="greeting-input"
          name="greeting"
          type="text"
          defaultValue={initialGreeting}
          required
        />
        <button className={styles.button} type="submit">
          Save
        </button>
      </form>
    </section>
  );
}
