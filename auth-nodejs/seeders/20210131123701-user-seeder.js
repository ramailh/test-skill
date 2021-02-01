'use strict';

module.exports = {
  up: async (queryInterface, Sequelize) => {
    /**
     * Add seed commands here.
     *
     * Example:
     * await queryInterface.bulkInsert('People', [{
     *   name: 'John Doe',
     *   isBetaMember: false
     * }], {});
    */

   queryInterface.bulkInsert('Users', [
     {
       name: 'jono',
       password: "testing1",
       phone: "081383838981",
       role: "admin",
       createdAt: new Date(),
       updatedAt: new Date()
      },
      {
        name: 'joni',
        password: "testing2",
        phone: "081383838982",
        role: "admin",
        createdAt: new Date(),
        updatedAt: new Date()
       },
       {
        name: 'jone',
        password: "testing3",
        phone: "081383838983",
        role: "admin",
        createdAt: new Date(),
        updatedAt: new Date()
       },
  ], {});
  },

  down: async (queryInterface, Sequelize) => {
    /**
     * Add commands to revert seed here.
     *
     * Example:
     * await queryInterface.bulkDelete('People', null, {});
     */
    await queryInterface.bulkDelete('Users', null, {});
  }
};
