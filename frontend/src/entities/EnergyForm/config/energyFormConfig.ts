export interface IFormInput {
	address: string;
	buildingType: string;
	roomsCount?: number;
	residentsCount?: number;
	totalArea?: number;
	consumption: {
		[key: string]: number;
	};
	file: string;
}
