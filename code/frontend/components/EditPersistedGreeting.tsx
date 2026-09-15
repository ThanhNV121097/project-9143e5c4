"use client";

import { FormEvent, useEffect, useRef, useState } from "react";
import { getMockGreeting, saveMockGreeting } from "../lib/mock/edit-persisted-greeting";
import styles from "./EditPersistedGreeting.module.css";

export type EditPersistedGreetingProps = {
  initialGreeting: string;
};

export function EditPersistedGreeting({ initialGreeting }: EditPersistedGreetingProps) {
  const currentGreeting = getMockGreeting().greeting;
  const [greeting, setGreeting] = useState(currentGreeting || initialGreeting);
  const [inputValue, setInputValue] = useState(currentGreeting || initialGreeting);
  const inputRef = useRef<HTMLInputElement>(null);

  function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();

    const nextGreeting = inputValue.trim();

    if (!nextGreeting) {
      inputRef.current?.focus();
      return;
    }

    const savedGreeting = saveMockGreeting(nextGreeting).greeting;

    setGreeting(savedGreeting);
    setInputValue(savedGreeting);
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
          ref={inputRef}
          className={styles.input}
          id="greeting-input"
          name="greeting"
          type="text"
          value={inputValue}
          required
          onChange={(event) => setInputValue(event.target.value)}
        />
        <button className={styles.button} type="submit">
          Save
        </button>
      </form>
    </section>
  );
}
