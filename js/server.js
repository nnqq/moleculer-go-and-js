const { ServiceBroker } = require('moleculer');

const broker = new ServiceBroker({
  logLevel: 'info',
  transporter: 'nats://localhost:4222',
});

broker.createService({
  name: 'jsMath',
  actions: {
    add(ctx) {
      return Number(ctx.params.a) + Number(ctx.params.b);
    },
  },
});

broker.start().then(() => {
  setTimeout(async () => {
    console.log('JS calling Go service goMath.add ...');

    const goResponse = await broker.call('goMath.add', {
      a: 1,
      b: 2,
    });

    console.log('JS=>Go. Response from Go service goMath.add =>', goResponse);
  }, 5000);
});
