import { Meta, Story, DecoratorFn } from '@storybook/react'

import { WebStory } from '../../../components/WebStory'

import { DotcomGettingStartedPage } from './DotcomGettingStartedPage'

const decorator: DecoratorFn = story => <div className="p-3 container">{story()}</div>

const config: Meta = {
    title: 'web/batches/DotcomGettingStartedPage',
    decorators: [decorator],
}

export default config

export const Overview: Story = () => <WebStory>{() => <DotcomGettingStartedPage />}</WebStory>

Overview.parameters = {
    chromatic: {
        disableSnapshot: false,
    },
}

Overview.storyName = 'Overview'
