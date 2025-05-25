import EnergyForm from "../entities/EnergyForm/UI/EnergyForm";
import TableData from "../entities/TableData/UI/TableData";

export default function App() {
	return (
		<div
			style={{
				display: "flex",
				flexDirection: "column",
				justifyContent: "center",
				alignItems: "center",
			}}
		>
			<EnergyForm />
			<TableData title={"Данные для отправки"} />
		</div>
	);
}
