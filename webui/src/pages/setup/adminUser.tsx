import React, {forwardRef, useCallback} from "react";
import Button from "react-bootstrap/Button";
import Card from "react-bootstrap/Card";
import Col from "react-bootstrap/Col";
import Form from "react-bootstrap/Form";
import Row from "react-bootstrap/Row";
import {Error} from "../../lib/components/controls";

interface AdminUserSetupProps {
    onSubmit: () => Promise<void>;
    setupError: Error;
    disabled: boolean;
}

export const AdminUserSetup = forwardRef<HTMLInputElement, AdminUserSetupProps>(({
    onSubmit,
    setupError,
    disabled,
}, inputRef) => {
    const submitHandler = useCallback((e) => {
        onSubmit();
        e.preventDefault();
    }, [onSubmit]);
    return (
        <Row>
            <Col md={{offset: 2, span: 8}}>
                <Card className="setup-widget">
                    <Card.Header>Initial Setup</Card.Header>
                    <Card.Body>
                        <Card.Text>
                            Please specify the name of the first admin account to create, or leave it as the default.
                        </Card.Text>
                        <Form onSubmit={submitHandler}>
                            <Form.Group controlId="user-name" className="mb-3">
                                <Form.Control type="text" value="admin" placeholder="Admin Username" ref={inputRef} autoFocus/>
                            </Form.Group>

                            {!!setupError && <Error error={setupError}/>}
                            <Button variant="primary" disabled={disabled} type="submit">Setup</Button>
                        </Form>
                    </Card.Body>
                </Card>
            </Col>
        </Row>
    );
});