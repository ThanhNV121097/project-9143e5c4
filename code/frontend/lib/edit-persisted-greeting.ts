export type GreetingResponse = {
  greeting: string;
};

const apiBase = process.env.NEXT_PUBLIC_API_URL ?? "/api";

export async function getGreeting(): Promise<GreetingResponse> {
  const response = await fetch(`${serverApiBase()}/v1/greeting`, { cache: "no-store" });

  if (!response.ok) {
    throw new Error("Failed to load greeting.");
  }

  return response.json();
}

export async function saveGreeting(greeting: string): Promise<GreetingResponse> {
  const response = await fetch(`${apiBase}/v1/greeting`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ greeting }),
  });

  if (!response.ok) {
    throw new Error("Failed to save greeting.");
  }

  return response.json();
}

function serverApiBase() {
  return typeof window === "undefined" ? process.env.API_ORIGIN ?? "http://backend:8080" : apiBase;
}
