import mockOrderingAPIs from "./ordering"
import HttpRequestMock from 'http-request-mock';

export default function mockAPI(){
    const mock = HttpRequestMock.setup()
        mockOrderingAPIs(mock)
}