"use client";

import { FormEvent, useState } from "react";
import styles from "./ShowPersistedGreeting.module.css";

type ShowPersistedGreetingProps = {
  initialGreeting: string;
};

export function ShowPersistedGreeting({ initialGreeting }: ShowPersistedGreetingProps) {
  const [greeting, setGreeting] = useState(initialGreeting);
  const [draft, setDraft] = useState(initialGreeting);

  function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    const nextGreeting = draft.trim();

    if (!nextGreeting) {
      return;
    }

    setGreeting(nextGreeting);
    setDraft(nextGreeting);
  }

  return (
    <section className={styles.section} aria-labelledby="greeting-heading">
      <h1 id="greeting-heading" className={styles.heading}>
        {greeting}
      </h1>
      <form className={styles.form} autoComplete="off" onSubmit={handleSubmit}>
        <label className={styles.label} htmlFor="greeting-input">
          Greeting
        </label>
        <input
          className={styles.input}
          id="greeting-input"
          name="greeting"
          type="text"
          value={draft}
          required
          onChange={(event) => setDraft(event.target.value)}
        />
        <button className={styles.button} type="submit">
          Save
        </button>
      </form>
    </section>
  );
}
