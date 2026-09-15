export type GreetingResponse = {
  greeting: string;
};

export async function getGreeting(): Promise<GreetingResponse> {
  return { greeting: "Hello, World!" };
}
