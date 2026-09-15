export type GreetingResponse = {
  greeting: string;
};

const storedGreeting: GreetingResponse = {
  greeting: "Hello, World!",
};

export function getMockGreeting(): GreetingResponse {
  return storedGreeting;
}
