import create from "zustand";
import { combine } from "zustand/middleware";

interface Store {
  user: {
    id: number;
    email: string;
    username: string;
  } | null;
}

export const useUserStore = create<Store>((set) => ({
  user: null,
  setUser: (user: Store["user"]) => set({ user }),
}));
