import { useStore } from "../../../app/store/store";
import { headersTable } from "../config/tableDataConfig";

interface ITableData {
	title: string;
}

const TableData = ({ title }: ITableData) => {
	const { state, setState } = useStore();
	const handleDelete = (accountId) => {
		setState((prev) => prev.filter((item) => item.accountId !== accountId));
	};

	return (
		<div>
			<p>{title}</p>
			<table>
				<tr>
					{headersTable.map(({ val }) => (
						<td>{val}</td>
					))}
				</tr>
				{state.length > 0 &&
					state.map(
						({
							accountId,
							address,
							roomsCount,
							residentsCount,
							totalArea,
							buildingType,
						}) => (
							<tr>
								<td>{accountId}</td>
								<td>{address}</td>
								<td>{roomsCount}</td>
								<td>{residentsCount}</td>
								<td>{totalArea}</td>
								<td>{buildingType}</td>
								<td>
									<button onClick={() => handleDelete(accountId)}>
										Удалить
									</button>
								</td>
							</tr>
						)
					)}
			</table>
		</div>
	);
};

export default TableData;
