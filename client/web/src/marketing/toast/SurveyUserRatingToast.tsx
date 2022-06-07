import React from 'react'

import { Button, Checkbox } from '@sourcegraph/wildcard'

import { SurveyRatingRadio } from '../components/SurveyRatingRadio'

import { Toast } from './Toast'

import styles from './SurveyUserRatingToast.module.scss'

export interface SurveyUserRatingToastProps {
    onChange?: (score: number) => void
    toggleErrorMessage: boolean
    onContinue: () => void
    onDismiss: () => void
    shouldPermanentlyDismiss: boolean
    toggleShouldPermanentlyDismiss: (value: boolean) => void
}

export const SurveyUserRatingToast: React.FunctionComponent<SurveyUserRatingToastProps> = ({
    onChange,
    toggleErrorMessage,
    onDismiss,
    onContinue,
    shouldPermanentlyDismiss,
    toggleShouldPermanentlyDismiss,
}) => (
    <Toast
        title="Tell us what you think"
        subtitle={
            <span id="survey-toast-scores">How likely is it that you would recommend Sourcegraph to a friend?</span>
        }
        cta={
            <>
                <SurveyRatingRadio ariaLabelledby="survey-toast-scores" onChange={onChange} />
                {toggleErrorMessage && (
                    <div className={styles.alertDanger} role="alert">
                        Please select a score between 0 to 10
                    </div>
                )}
            </>
        }
        footer={
            <div className="d-flex align-items-center justify-content-between">
                <Checkbox
                    id="survey-toast-refuse"
                    label={<span className={styles.checkboxLabel}>Don't show this again</span>}
                    checked={shouldPermanentlyDismiss}
                    onChange={event => toggleShouldPermanentlyDismiss(event.target.checked)}
                />
                <Button variant="secondary" size="sm" onClick={onContinue}>
                    Continue
                </Button>
            </div>
        }
        onDismiss={onDismiss}
    />
)
