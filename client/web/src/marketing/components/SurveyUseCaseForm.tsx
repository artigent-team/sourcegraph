import React, { useEffect, useState } from 'react'

import classNames from 'classnames'

import { FlexTextArea, H4, Input } from '@sourcegraph/wildcard'

import { AuthenticatedUser } from '../../auth'
import { SurveyUseCase } from '../../graphql-operations'

import { SurveyUseCaseCheckbox } from './SurveyUseCaseCheckbox'

import styles from './SurveyUseCaseForm.module.scss'

export const OPTIONS = [
    {
        id: SurveyUseCase.UNDERSTAND_NEW_CODE,
        labelValue: 'Understand a new codebase',
    },
    {
        id: SurveyUseCase.FIX_SECURITY_VULNERABILITIES,
        labelValue: 'Fix security vulnerabilities',
    },
    {
        id: SurveyUseCase.REUSE_CODE,
        labelValue: 'Reuse code',
    },
    {
        id: SurveyUseCase.RESPOND_TO_INCIDENTS,
        labelValue: 'Respond to incidents',
    },
    {
        id: SurveyUseCase.IMPROVE_CODE_QUALITY,
        labelValue: 'Improve code quality',
    },
] as const

interface SurveyUseCaseFormProps {
    onChangeUseCases: (useCases: SurveyUseCase[]) => void
    onChangeOtherUseCase: (others: string) => void
    onChangeAdditionalInformation: (additionalInformation: string) => void
    onChangeEmail: (email: string) => void
    formLabelClassName?: string
    additionalInformation: string
    otherUseCase: string
    email: string
    className?: string
    title: string
    authenticatedUser?: AuthenticatedUser | null
}

export const SurveyUseCaseForm: React.FunctionComponent<SurveyUseCaseFormProps> = ({
    onChangeAdditionalInformation,
    onChangeOtherUseCase,
    onChangeUseCases,
    onChangeEmail,
    formLabelClassName,
    additionalInformation,
    otherUseCase,
    email,
    className,
    title,
    authenticatedUser,
}) => {
    const [useCases, setUseCases] = useState<SurveyUseCase[]>([])
    const [showOtherInput, setShowOtherInput] = useState<boolean>(false)

    const handleSelectUseCase = (value: SurveyUseCase): void => {
        if (useCases.includes(value)) {
            setUseCases(current => current.filter(instance => instance !== value))
            return
        }
        setUseCases(current => [...current, value])
    }

    useEffect(() => {
        onChangeUseCases(useCases)
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [useCases])

    return (
        <div className={classNames('mb-2', className)}>
            <H4 id="usecase-group" className={classNames('d-flex', styles.title, formLabelClassName)}>
                {title}
            </H4>
            <fieldset className={styles.checkWrap} aria-labelledby="usecase-group">
                {OPTIONS.map(({ id, labelValue }) => (
                    <SurveyUseCaseCheckbox
                        label={labelValue}
                        onChange={() => handleSelectUseCase(id)}
                        key={id}
                        id={id}
                        checked={useCases.includes(id)}
                    />
                ))}
                <SurveyUseCaseCheckbox
                    label="Other"
                    onChange={() => setShowOtherInput(!showOtherInput)}
                    id="survey_checkbox_other"
                    checked={showOtherInput}
                />
            </fieldset>
            {showOtherInput && (
                <FlexTextArea
                    containerClassName="mt-3"
                    label={
                        <span className={classNames(styles.textareaLabel, formLabelClassName)}>
                            What else are you using Sourcegraph to do?
                        </span>
                    }
                    onChange={event => onChangeOtherUseCase(event.target.value)}
                    value={otherUseCase}
                />
            )}
            <FlexTextArea
                containerClassName="mt-3"
                label={
                    <span className={classNames(styles.textareaLabel, formLabelClassName)}>
                        Anything else you would like to share with us?
                    </span>
                }
                onChange={event => onChangeAdditionalInformation(event.target.value)}
                value={additionalInformation}
            />
            {!authenticatedUser && (
                <Input
                    className="mt-3"
                    label={
                        <span className={classNames(styles.textareaLabel, formLabelClassName)}>
                            What is your email?
                        </span>
                    }
                    onChange={event => onChangeEmail(event.target.value)}
                    value={email}
                />
            )}
        </div>
    )
}
