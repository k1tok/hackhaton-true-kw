import { createContext, useContext } from "react";
import type { IFormInput } from "../../entities/EnergyForm/config/energyFormConfig";

export const StoreContext = createContext({
	state: [] as IFormInput[],
	setState: () => {},
});

export const useStore = () => useContext(StoreContext);
