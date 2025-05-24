import { useForm, type SubmitHandler } from "react-hook-form";
import FieldsEnergyForm from "../../../widgets/FieldsEnergyForm/UI/FieldsEnergyForm";
import type { IFormInput } from "../config/energyFormConfig";
import styles from "../styles/EnergyForm.module.css";
import ButtonSubmitForm from "../../../shared/ButtonSubmitForm/UI/ButtonSubmitForm";
import { useStore } from "../../../app/store/store";
import { parseJson } from "../../../tools/parseJson";

const EnergyForm = () => {
	const { state, setState } = useStore();

	const addData = (data) => {
		setState((prev) => [...prev, ...data]);
	};

	const { register, handleSubmit, reset } = useForm<IFormInput>();

	const onSubmit: SubmitHandler<IFormInput> = async (data) => {
		const jsonData = await parseJson(data);
		const finalData = [data, ...jsonData];

		addData(finalData);
		console.log(data);
		reset();
	};

	return (
		<form onSubmit={handleSubmit(onSubmit)} className={styles.form}>
			<div className={styles.title}>Объект проверки</div>
			<FieldsEnergyForm register={register} />
			<ButtonSubmitForm title="Добавить данные" />
		</form>
	);
};

export default EnergyForm;
