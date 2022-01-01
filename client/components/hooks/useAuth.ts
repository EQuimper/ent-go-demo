import { useQuery } from "react-query";

async function getUser() {
  const res = await fetch("/api/me", {
    credentials: "include",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
  });

  if (res.status >= 400) {
    return null;
  }

  const { data } = await res.json();

  return data;
}

export const useAuth = () => {
  const { data, isLoading } = useQuery<{
    id: number;
    username: string;
    email: string;
  }>("me", getUser);

  return {
    data,
    isLoading,
    isLogged: data !== null,
  };
};
