/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package org.apache.beam.dsls.sql;

import org.apache.beam.dsls.sql.schema.BeamSqlRow;
import org.apache.beam.sdk.testing.PAssert;
import org.apache.beam.sdk.values.PCollection;
import org.apache.beam.sdk.values.PCollectionTuple;
import org.apache.beam.sdk.values.TupleTag;
import org.junit.Test;

/**
 * Tests for WHERE queries.
 */
public class BeamSqlDslFilterTest extends BeamSqlDslBase {
  /**
   * single filter.
   */
  @Test
  public void testSingleFilter() throws Exception {
    String sql = "SELECT * FROM PCOLLECTION WHERE f_int = 1";

    PCollection<BeamSqlRow> result =
        inputA1.apply("testSingleFilter", BeamSql.simpleQuery(sql));

    PAssert.that(result).containsInAnyOrder(recordsInTableA.get(0));

    pipeline.run().waitUntilFinish();
  }

  /**
   * composite filters.
   */
  @Test
  public void testCompositeFilter() throws Exception {
    String sql = "SELECT * FROM TABLE_A"
        + " WHERE f_int > 1 AND (f_long < 3000 OR f_string = 'string_row3')";

    PCollection<BeamSqlRow> result =
        PCollectionTuple.of(new TupleTag<BeamSqlRow>("TABLE_A"), inputA1)
        .apply("testCompositeFilter", BeamSql.query(sql));

    PAssert.that(result).containsInAnyOrder(recordsInTableA.get(1), recordsInTableA.get(2));

    pipeline.run().waitUntilFinish();
  }

  /**
   * nothing return with filters.
   */
  @Test
  public void testNoReturnFilter() throws Exception {
    String sql = "SELECT * FROM TABLE_A WHERE f_int < 1";

    PCollection<BeamSqlRow> result =
        PCollectionTuple.of(new TupleTag<BeamSqlRow>("TABLE_A"), inputA1)
        .apply("testNoReturnFilter", BeamSql.query(sql));

    PAssert.that(result).empty();

    pipeline.run().waitUntilFinish();
  }

  @Test
  public void testFromInvalidTableName1() throws Exception {
    exceptions.expect(IllegalStateException.class);
    exceptions.expectMessage("Object 'TABLE_B' not found");
    pipeline.enableAbandonedNodeEnforcement(false);

    String sql = "SELECT * FROM TABLE_B WHERE f_int < 1";

    PCollection<BeamSqlRow> result =
        PCollectionTuple.of(new TupleTag<BeamSqlRow>("TABLE_A"), inputA1)
        .apply("testFromInvalidTableName1", BeamSql.query(sql));

    pipeline.run().waitUntilFinish();
  }

  @Test
  public void testFromInvalidTableName2() throws Exception {
    exceptions.expect(IllegalStateException.class);
    exceptions.expectMessage("Use fixed table name PCOLLECTION");
    pipeline.enableAbandonedNodeEnforcement(false);

    String sql = "SELECT * FROM PCOLLECTION_NA";

    PCollection<BeamSqlRow> result = inputA1.apply(BeamSql.simpleQuery(sql));

    pipeline.run().waitUntilFinish();
  }

  @Test
  public void testInvalidFilter() throws Exception {
    exceptions.expect(IllegalStateException.class);
    exceptions.expectMessage("Column 'f_int_na' not found in any table");
    pipeline.enableAbandonedNodeEnforcement(false);

    String sql = "SELECT * FROM PCOLLECTION WHERE f_int_na = 0";

    PCollection<BeamSqlRow> result = inputA1.apply(BeamSql.simpleQuery(sql));

    pipeline.run().waitUntilFinish();
  }
}