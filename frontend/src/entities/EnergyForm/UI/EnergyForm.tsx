import { useForm, type SubmitHandler } from "react-hook-form";
import FieldsEnergyForm from "../../../widgets/FieldsEnergyForm/UI/FieldsEnergyForm";
import type { IFormInput } from "../config/energyFormConfig";
import styles from "../styles/EnergyForm.module.css";
import ButtonSubmitForm from "../../../shared/ButtonSubmitForm/UI/ButtonSubmitForm";

const EnergyForm = () => {
	const { register, handleSubmit } = useForm<IFormInput>();
	const onSubmit: SubmitHandler<IFormInput> = (data) => {
		console.log(data);
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
