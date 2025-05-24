import EnergyForm from "../entities/EnergyForm/UI/EnergyForm";

export default function App() {
	return (
		<div>
			<EnergyForm />
			{/* <div>
				<p>Данные</p>
				<table>
					<tr>
						<td>Номер обращения</td>
						<td>Адрес</td>
						<td>Тип строения</td>
						<td>Количество комнат</td>
						<td>Количество жильцов</td>
						<td>Площадь помещения</td>
						<td>Потребление энергии</td>
					</tr>
					{state.map(
						(
							{
								address,
								buildingType,
								roomsCount,
								residentsCount,
								totalArea,
								consumption,
							}: IFormInput,
							index
						) => {
							return (
								<tr>
									<td>{index + 1000}</td>
									<td>{address}</td>
									<td>{buildingType}</td>
									<td>{roomsCount}</td>
									<td>{residentsCount}</td>
									<td>{totalArea}</td>
									{consumption.map((month) => {
										return (
											<td>
												<div>Месяц</div>
												<div>{month}</div>
											</td>
										);
									})}
								</tr>
							);
						}
					)}
				</table>
			</div> */}
		</div>
	);
}
