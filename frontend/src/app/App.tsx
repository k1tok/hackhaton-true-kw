import { useState } from "react";
import MainPage from "../pages/MainPage";
import { StoreContext } from "./store/store";
import "./styles/App.css";
import { type IFormInput } from "../entities/EnergyForm/config/energyFormConfig";

function App() {
	const [state, setState] = useState<IFormInput[]>([]);

	return (
		<StoreContext.Provider value={{ state, setState }}>
			<MainPage />
		</StoreContext.Provider>
	);
}

export default App;
