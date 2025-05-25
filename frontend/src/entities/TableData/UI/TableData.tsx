import { useEffect, useState } from "react";
import { useStore } from "../../../app/store/store";
import { headersTable, renderConsumption } from "../config/tableDataConfig";
import styles from "./../style/TableData.module.css";
import { fetchData } from "../../../tools/fetchData";

interface ITableData {
	title: string;
}

const TableData = ({ title }: ITableData) => {
	const { state, setState } = useStore();
	const handleDelete = (accountId) => {
		setState((prev) => prev.filter((item) => item.accountId !== accountId));
	};

	const handleSend = async () => {
		const req = await fetchData("http://127.0.0.1/check", state);
		const resp = await req.json();
		return req;
	};

	return (
		<div className={styles.container}>
			{state.length > 0 && (
				<>
					<p className={styles.title}>{title}</p>
					<table className={styles.table}>
						<thead>
							<tr>
								{headersTable.map(({ val }) => (
									<th key={val}>{val}</th>
								))}
							</tr>
						</thead>
						<tbody>
							{state.length > 0 &&
								state.map(
									({
										accountId,
										address,
										roomsCount,
										residentsCount,
										totalArea,
										buildingType,
										consumption,
									}) => (
										<tr key={accountId}>
											<td>{accountId}</td>
											<td>{address}</td>
											<td>{roomsCount}</td>
											<td>{residentsCount}</td>
											<td>{totalArea}</td>
											<td>{buildingType}</td>
											<td className={styles.consumptionCell}>
												{renderConsumption(consumption)}
											</td>
											<td>
												<button
													className={styles.deleteButton}
													onClick={() => handleDelete(accountId)}
												>
													Удалить
												</button>
											</td>
										</tr>
									)
								)}
						</tbody>
					</table>
					{state.length > 0 && (
						<button onClick={handleSend}>Отправить данные</button>
					)}
				</>
			)}
		</div>
	);
};

export default TableData;
