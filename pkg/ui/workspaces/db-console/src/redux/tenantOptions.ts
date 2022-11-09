// Copyright 2022 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.
export const selectTenantsFromCookies = (): string[] => {
  const sessionCookieStr = document.cookie
    .split(";")
    .filter(row => row.startsWith("session="))[0];
  return sessionCookieStr
    ? sessionCookieStr
        .replace("session=", "")
        .split(/[,&]/g)
        .filter((_, idx) => idx % 2 == 1)
    : [];
};

export const selectCurrentTenantIDFromCookies = (): string | null => {
  const tenantCookieStr = document.cookie
    .split(";")
    .filter(row => row.startsWith("tenant="))[0];
  return tenantCookieStr ? tenantCookieStr.replace("tenant=", "") : null;
};
