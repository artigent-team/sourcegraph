// This file is a fork from SearchBox.tsx and contains JetBrains specific UI changes
/* eslint-disable no-restricted-imports */

import React, { useCallback, useState } from 'react'

import classNames from 'classnames'
import * as Monaco from 'monaco-editor'

import { QueryState, SearchContextInputProps, SubmitSearchProps } from '@sourcegraph/search'
import { DEFAULT_MONACO_OPTIONS } from '@sourcegraph/search-ui'
import {
    IEditor,
    LazyMonacoQueryInput,
    LazyMonacoQueryInputProps,
} from '@sourcegraph/search-ui/src/input/LazyMonacoQueryInput'
import { SearchContextDropdown } from '@sourcegraph/search-ui/src/input/SearchContextDropdown'
import { AuthenticatedUser } from '@sourcegraph/shared/src/auth'
import { KeyboardShortcut } from '@sourcegraph/shared/src/keyboardShortcuts'
import { PlatformContextProps } from '@sourcegraph/shared/src/platform/context'
import { fetchStreamSuggestions as defaultFetchStreamSuggestions } from '@sourcegraph/shared/src/search/suggestions'
import { TelemetryProps } from '@sourcegraph/shared/src/telemetry/telemetryService'
import { ThemeProps } from '@sourcegraph/shared/src/theme'

import { Search } from '../jetbrains-icons/Search'

import { JetBrainsToggles, JetBrainsTogglesProps } from './JetBrainsToggles'

import styles from './JetBrainsSearchBox.module.scss'

export interface JetBrainsSearchBoxProps
    extends Omit<JetBrainsTogglesProps, 'navbarSearchQuery' | 'submitSearch' | 'clearSearch'>,
        ThemeProps,
        SearchContextInputProps,
        TelemetryProps,
        PlatformContextProps<'requestGraphQL'>,
        Pick<LazyMonacoQueryInputProps, 'editorComponent'> {
    authenticatedUser: AuthenticatedUser | null
    isSourcegraphDotCom: boolean // significant for query suggestions
    showSearchContext: boolean
    showSearchContextManagement: boolean
    queryState: QueryState
    onChange: (newState: QueryState) => void
    onSubmit: () => void
    submitSearchOnSearchContextChange?: SubmitSearchProps['submitSearch']
    submitSearchOnToggle?: SubmitSearchProps['submitSearch']
    onFocus?: () => void
    fetchStreamSuggestions?: typeof defaultFetchStreamSuggestions // Alternate implementation is used in the VS Code extension.
    onCompletionItemSelected?: () => void
    onSuggestionsInitialized?: (actions: { trigger: () => void }) => void
    autoFocus?: boolean
    keyboardShortcutForFocus?: KeyboardShortcut
    className?: string
    containerClassName?: string

    /** Whether globbing is enabled for filters. */
    globbing: boolean

    /** Whether comments are parsed and highlighted */
    interpretComments?: boolean

    /** Don't show search help button */
    hideHelpButton?: boolean

    onHandleFuzzyFinder?: React.Dispatch<React.SetStateAction<boolean>>

    /** Set in JSContext only available to the web app. */
    isExternalServicesUserModeAll?: boolean

    /** Called with the underlying editor instance on creation. */
    onEditorCreated?: (editor: IEditor) => void
}

const MONACO_OPTIONS: Monaco.editor.IStandaloneEditorConstructionOptions = {
    ...DEFAULT_MONACO_OPTIONS,

    // Reset to default value to avoid suggestion box shrinking.
    // Related issue: https://github.com/microsoft/monaco-editor/issues/3147
    suggestLineHeight: undefined,
}

export const JetBrainsSearchBox: React.FunctionComponent<React.PropsWithChildren<JetBrainsSearchBoxProps>> = props => {
    const { queryState, onEditorCreated: onEditorCreatedCallback, onChange } = props

    const [editor, setEditor] = useState<IEditor>()
    const focusEditor = useCallback(() => editor?.focus(), [editor])

    const onEditorCreated = useCallback(
        (editor: IEditor) => {
            setEditor(editor)
            onEditorCreatedCallback?.(editor)
        },
        [onEditorCreatedCallback]
    )

    const clearSearch = (): void => {
        onChange({ ...queryState, query: '' })
        focusEditor()
    }

    return (
        <div className={classNames(styles.searchBox, props.containerClassName)}>
            <div
                className={classNames(
                    styles.searchBoxBackgroundContainer,
                    props.className,
                    'flex-shrink-past-contents'
                )}
            >
                {props.searchContextsEnabled && props.showSearchContext && (
                    <>
                        <SearchContextDropdown
                            /* eslint-disable-next-line no-restricted-syntax */
                            {...props}
                            query={queryState.query}
                            submitSearch={props.submitSearchOnSearchContextChange}
                            className={styles.searchBoxContextDropdown}
                            menuClassName={styles.searchBoxContextMenu}
                            onEscapeMenuClose={focusEditor}
                        />
                        <div className={styles.searchBoxSeparator} />
                    </>
                )}
                {/*
                    To fix Rule: "region" (All page content should be contained by landmarks)
                    Added role attribute to the following element to satisfy the rule.
                */}
                <div className={classNames(styles.searchBoxFocusContainer, 'flex-shrink-past-contents')} role="search">
                    <div className={styles.searchBoxFocusContainerIcon}>
                        <Search />
                    </div>
                    <LazyMonacoQueryInput
                        /* eslint-disable-next-line no-restricted-syntax */
                        {...props}
                        onHandleFuzzyFinder={props.onHandleFuzzyFinder}
                        className={styles.searchBoxInput}
                        onEditorCreated={onEditorCreated}
                        placeholder="Enter search query..."
                        editorOptions={MONACO_OPTIONS}
                    />
                    <JetBrainsToggles
                        patternType={props.patternType}
                        setPatternType={props.setPatternType}
                        caseSensitive={props.caseSensitive}
                        setCaseSensitivity={props.setCaseSensitivity}
                        settingsCascade={props.settingsCascade}
                        submitSearch={props.submitSearchOnToggle}
                        navbarSearchQuery={queryState.query}
                        className={styles.searchBoxToggles}
                        showCopyQueryButton={props.showCopyQueryButton}
                        structuralSearchDisabled={props.structuralSearchDisabled}
                        selectedSearchContextSpec={props.selectedSearchContextSpec}
                        clearSearch={clearSearch}
                    />
                </div>
            </div>
        </div>
    )
}
