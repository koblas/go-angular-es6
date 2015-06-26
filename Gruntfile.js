module.exports = function (grunt) {
   grunt.initConfig({
      browserify: {
         dist: {
            options: {
               transform: [
                  ["babelify", {
                     loose: "all",
                     sourceMap: true,
                  }], "partialify"
               ],
               extensions: ['.es6'],
               debug: true
            },
            files: {
               "./static/app.js": ["./client/app/main.es6"]
            }
         }
      },
      copy: {
          all: {
              // This copies all the html and css into the dist/ folder
              expand: true,
              cwd: 'client/',
              src: ['**/*.css', 'components/**', 'bootstrap/**', 'app.html'],
              dest: 'static/',
          },
      },
      watch: {
         grunt: {
            files: ["Gruntfile.js"],
            options: {
                reload: true
            }
         },
         app: {
            files: ["./client/app/**", "./client/partial/**"],
            tasks: ["browserify"]
         },
         files: {
            files: ['./client/**/*.css', './client/components/**', './client/bootstrap/**', './client/app.html'],
            tasks: ["copy"]
         }
      }
   });
 
   grunt.loadNpmTasks("grunt-browserify");
   grunt.loadNpmTasks("grunt-contrib-watch");
   grunt.loadNpmTasks('grunt-contrib-copy');
 
   grunt.registerTask("default", ["watch"]);
   grunt.registerTask("build", ["browserify", "copy"]);
};
