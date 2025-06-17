interface UserSession {
  token: string;
  userId: number;
  role: "customer" | "seller";
  email: string;
}

export const setUserSession = (session: UserSession) => {
  localStorage.setItem("userSession", JSON.stringify(session));
};

export const getUserSession = (): UserSession | null => {
  const session = localStorage.getItem("userSession");
  return session ? JSON.parse(session) : null;
};

export const clearUserSession = () => {
  localStorage.removeItem("userSession");
};

export const isAuthenticated = (): boolean => {
  return !!getUserSession();
};

export const isSeller = (): boolean => {
  const session = getUserSession();
  return session?.role === "seller";
};

export const getUserId = (): number | null => {
  const session = getUserSession();
  return session?.userId || null;
};

export const getToken = (): string | null => {
  const session = getUserSession();
  return session?.token || null;
}; 