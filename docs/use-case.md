
![UseCase](https://g.gravizo.com/svg?
  @startuml;
  left to right direction;
  skinparam packageStyle rectangle;
  actor user;
  actor devloper;
  actor store_admin;
  actor yak_developer;
  rectangle host {;
    %28edit config%29 -> %28service connection%29;
    %28fetch recipes%29 -> %28edit config%29;
    %28recipe devices%29 -> %28fetch recipes%29;
    rectangle yak_exec {;
       %28fetch recipes%29 -> %28install application%29;
       user .> %28install application%29;
       %28install application%29 -> %28upgrade%29;
       %28upgrade%29 .> %28trigger hook%29;
       %28install application%29 .> %28trigger hook%29;
       %28edit config%29 .> %28trigger hook%29;
       %28service connection%29 .> %28trigger hook%29;
    };
  };
  rectangle store {;
    user -> %28register host%29;
    devloper -> %28manage recipe%29;
    store_admin -> %28manage alias%29;
    yak_developer -> %28define backend%29;
    rectangle host_recipe {;
       %28manage recipe%29 -> %28recipe users%29;
       %28manage recipe%29 -> %28recipe devices%29;
       %28manage recipe%29 -> %28recipe configs%29;
       %28register host%29 -> %28recipe configs%29;
       %28register host%29 -> %28recipe users%29;
       %28register host%29 -> %28recipe devices%29;
    };
  };
  @enduml
)

![components](https://g.gravizo.com/svg?
  @startuml;
  node "host" {;
    frame host_data {;
      folder host_recipes {;
        [recipe_1];
        [recipe_2];
        [recipe_3];
      };
      database "sqlite" {;
        [host_info];
      };
    };
    frame "client" {;
        [web_client];
        [cli];
    };
    sqlite - yakd;
    yakd - host_recipes;
    web_client .> yakd;
    cli .> yakd;
  };
  ;
  cloud "store" {;
    frame data {;
      folder all_recipes {;
        [recipe_x];
        [recipe_y];
        [recipe_z];
      };
      database psql {;
        [hosts];
        [users];
        [recipes];
      };
    };
    goyak.io - psql;
  };
  goyak.io - yakd;
  web_client -> goyak.io;
  cli -> goyak.io;
  ;
  @enduml
)
