
import { createContext, ReactNode, useContext, useState } from "react";

interface ContextProps {
  profile?:any,
  setProfile?: any,
}
const StateContext = createContext<ContextProps>({});

interface ContextProviderProps {
  children?: ReactNode;
}
export const ContextProvider = ({ children }: ContextProviderProps) => {
  const [profile, setProfile] = useState(null)

  return (
    <StateContext.Provider
      value={{
        profile,
        setProfile
      }}
    >
      {children}
    </StateContext.Provider>
  );
};

export const useStateContext = () => useContext(StateContext);

