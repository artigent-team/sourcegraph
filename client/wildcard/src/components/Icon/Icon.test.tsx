import { render } from '@testing-library/react'

import { SourcegraphIcon } from '../SourcegraphIcon'

import { Icon } from './Icon'

describe('Icon', () => {
    it('renders a simple inline icon correctly', () => {
        const { asFragment } = render(<Icon as={SourcegraphIcon} aria-hidden={true} />)
        expect(asFragment()).toMatchSnapshot()
    })
    it('renders a medium icon correctly', () => {
        const { asFragment } = render(<Icon as={SourcegraphIcon} size="md" aria-hidden={true} />)
        expect(asFragment()).toMatchSnapshot()
    })
})
