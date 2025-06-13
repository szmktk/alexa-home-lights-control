const tplink = require('tplink-cloud-api');

tplink.login('<your-email>', '<your-password>').then(async () => {
  const devices = await tplink.getDeviceList();
  const device = tplink.getHS100(devices[0].deviceId); // or KL110

  await device.setPowerState(true); // turn on
  await device.setBrightness(50);   // set brightness to 50%
});
