export type GreetingResponse = {
  greeting: string;
};

export async function getGreeting(): Promise<GreetingResponse> {
  const apiBase = process.env.API_ORIGIN ?? "http://backend:8080";
  const response = await fetch(`${apiBase}/v1/greeting`, { cache: "no-store" });

  if (!response.ok) {
    throw new Error("Failed to load greeting.");
  }

  return response.json() as Promise<GreetingResponse>;
}
