export type GreetingResponse = {
  greeting: string;
};

const storageKey = "hello-world-acceptance-8:greeting";

const storedGreeting: GreetingResponse = {
  greeting: "Hello, World!",
};

export function getMockGreeting(): GreetingResponse {
  if (typeof window === "undefined") {
    return storedGreeting;
  }

  return {
    greeting: window.localStorage.getItem(storageKey) ?? storedGreeting.greeting,
  };
}

export function saveMockGreeting(greeting: string): GreetingResponse {
  storedGreeting.greeting = greeting;

  if (typeof window !== "undefined") {
    window.localStorage.setItem(storageKey, greeting);
  }

  return storedGreeting;
}
