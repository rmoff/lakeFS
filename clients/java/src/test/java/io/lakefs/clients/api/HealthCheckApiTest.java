/*
 * lakeFS API
 * lakeFS HTTP API
 *
 * The version of the OpenAPI document: 0.1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package io.lakefs.clients.api;

import io.lakefs.clients.api.ApiException;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for HealthCheckApi
 */
@Ignore
public class HealthCheckApiTest {

    private final HealthCheckApi api = new HealthCheckApi();

    
    /**
     * 
     *
     * check that the API server is up and running
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void healthCheckTest() throws ApiException {
                api.healthCheck();
        // TODO: test validations
    }
    
}
