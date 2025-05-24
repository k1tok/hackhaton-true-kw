import { useEffect, useState } from "react";
import { useStore } from "../../../app/store/store";
import { headersTable, renderConsumption } from "../config/tableDataConfig";
import styles from "./../style/TableData.module.css";

interface ITableData {
	title: string;
}

const TableData = ({ title }: ITableData) => {
	const { state, setState } = useStore();
	const handleDelete = (accountId) => {
		setState((prev) => prev.filter((item) => item.accountId !== accountId));
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
				</>
			)}
		</div>
	);
};

export default TableData;
