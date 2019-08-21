# -*- mode: ruby -*-
# vi: set ft=ruby :

# Windows users must manually set a shell variable for X11 forwarding
if ARGV[0] =~ /^up|provision$/i and not ARGV.include?("--no-provision")  then
  if Vagrant::Util::Platform.windows? then
    STDERR.puts "Windows detected. Remember to set DISPLAY to enable X11 forwarding."
  end
end

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.

Vagrant.configure(2) do |config|

  # 64 bit Ubuntu Vagrant Box
  #config.vm.box = "ubuntu/trusty64"
  #config.vm.box = "ubuntu/xenial64"
  config.vm.box = "ubuntu/bionic64"

  ## Configure hostname and port forwarding
  config.vm.hostname = "cos316"
  config.ssh.forward_x11 = true
  config.vm.network "forwarded_port", guest: 8888, host: 8888
  # Tinode server
  config.vm.network "forwarded_port", guest: 6060, host: 6060
  # Tegola server
  config.vm.network "forwarded_port", guest: 7070, host: 7070

  vagrant_root = File.dirname(__FILE__)

  # Synced Assignment Directories
  # Replace "<path>" with the actual path from the vagrantfile to your cloned
  # assignment directory, and uncomment the line for the appropriate assignment:

  # config.vm.synced_folder "<path to A1 repo>", "/vagrant/assignments/assignment1"
  # config.vm.synced_folder "<path to A2 repo>", "/vagrant/assignments/assignment2"
  # config.vm.synced_folder "<path to A3 repo>", "/vagrant/assignments/assignment3"
  # config.vm.synced_folder "<path to A4 repo>", "/vagrant/assignments/assignment4"
  # config.vm.synced_folder "<path to A5 repo>", "/vagrant/assignments/assignment5"
  # config.vm.synced_folder "<path to A6 repo>", "/vagrant/assignments/assignment6"
  # config.vm.synced_folder "<path to A7 repo>", "/vagrant/assignments/assignment7"



  # Emacs settings
  config.vm.provision "file", source: "#{vagrant_root}/config_files/dot_emacs", destination: "~/.emacs"

  # Jupyter notebook settings
  config.vm.provision "file", source: "#{vagrant_root}/config_files/jupyter_notebook_config.py", destination: "~/.jupyter/jupyter_notebook_config.py"

  # Flottbot config
  config.vm.provision "file", source: "#{vagrant_root}/config_files/flottbot/config", destination: "~/flottbot/config"

  # Tegola config
  config.vm.provision "file", source: "#{vagrant_root}/config_files/tegola", destination: "~/tegola"

  ## Provisioning
  config.vm.provision "shell", inline: <<-SHELL

    # Ask politely for apt-get not to expect any user input (hide "stdin" error msgs)
    export DEBIAN_FRONTEND=noninteractive

    # Ask less politely.
    debconf="/etc/apt/apt.conf.d/70debconf"
    sudo sed -i.bak "s|^DPkg::Pre-Install|//&|" $debconf # comment old preconfigure
    echo 'DPkg::Pre-Install-Pkgs {"/usr/sbin/dpkg-preconfigure --frontend=noninteractive --terse --apt || true";};' >> $debconf

    # disp "APT Updates"
    sudo add-apt-repository -y ppa:gophers/archive 2> /dev/null # For Go 1.11
    sudo apt-get -yq -o=Dpkg::Use-Pty=0 update
    sudo apt-get -yq -o=Dpkg::Use-Pty=0 upgrade

    # Alias apt-get installer with options to suppress useless output, progress bars
    install="sudo -E apt-get -yq -o=Dpkg::Use-Pty=0 install"

    # install programs needed for setup + testing
    $install toilet
    disp() { toilet -f mono9 $1; }
    disp "Welcome to COS316"

    $install jq
    disp "Emacs 25"
    $install emacs-nox

    disp "Go v1.11"
    $install golang-1.11-go
    echo "export PATH=${PATH}:/usr/lib/go-1.11/bin:/go/bin" >> /home/vagrant/.profile
    echo "export GOPATH='/go'" >> /home/vagrant/.profile
    # echo "export GO111MODULE=auto" >> /home/vagrant/.profile
    source /home/vagrant/.profile # Make sure go environment variables are set

    disp "MySQL 5.7"
    $install mysql-server-5.7

    disp "Tinode"
    $install git

    printf "Downloading tinode server... This may take some time."
    go get -tags mysql github.com/tinode/chat/server && go install -tags mysql github.com/tinode/chat/server && echo "tinode server installed"
    go get -tags mysql github.com/tinode/chat/tinode-db && go install -tags mysql github.com/tinode/chat/tinode-db && echo "tinode-db installed"
    printf "Download complete." # "Go installed the following packages:"
    # go list ... | grep tinode

    # Update tinode SQL password
    TINODE_PATH=/go/src/github.com/tinode/chat
    TINODE_STATIC=/home/vagrant/tinode/example-react-js
    echo "export TINODE_PATH=${TINODE_PATH}" >> /home/vagrant/.profile # Debugging convenience
    sudo sed -i.bak "s|root@|root:${MYSQL_PWD}@|" ${TINODE_PATH}/**/tinode.conf

    # Tinode expects these templates to be in /go/bin/templ, rather than /go/src/...
    # TODO Find the root cause of this -- it will likely cause problems later
    if [ ! -d $TINODE_STATIC ]; then
      mkdir /go/bin/templ
      cp $TINODE_PATH/server/templ/* /go/bin/templ
    fi
    printf "Done downloading tinode server"

    # Set up the static web files required by tinode
    $install unzip
    tinode-db -config=${TINODE_PATH}/tinode-db/tinode.conf -data=${TINODE_PATH}/tinode-db/data.json
    if [ ! -d $TINODE_STATIC ]; then
      echo "Downloading tinode-chat static files..."
      mkdir -p $TINODE_STATIC && cd $TINODE_STATIC
      sudo wget -q -O webapp-master.zip https://github.com/tinode/webapp/archive/master.zip
      sudo wget -q -O tinode-js-master.zip https://github.com/tinode/tinode-js/archive/master.zip
      echo "Downloaded static tinode files to $(pwd) ..."
      sudo unzip webapp-master.zip
      sudo unzip tinode-js-master.zip tinode-js-master/src/tinode.js
      # Now copy contents of unzipped folder into single dir where they are expected
      sudo cp -r webapp-master/* .
      sudo cp -r tinode-js-master/src/tinode.js .
      # cd elsewhere if desired
    else
      echo "Found existing tinode-chat files. Skipping..."
    fi

    # Install flottbot chatbot
    disp "Flottbot"
    cd /home/vagrant/flottbot
    sudo wget -q -O flottbot.tgz https://github.com/target/flottbot/releases/download/0.2.0/flottbot-linux-amd64.tgz
    tar xzf flottbot.tgz
    echo "Flottbot installed successfully."

    # disp "PostgreSQL"
    $install postgresql-10
    $install postgresql-10-postgis-2.4
    $install postgresql-10-postgis-2.4-scripts

    # # Install Tegola
    TEGOLA_PATH=/home/vagrant/tegola
    if [ ! -f $TEGOLA_PATH/tegola ]; then

      cd $TEGOLA_PATH
      wget -q -O tegola.zip https://github.com/go-spatial/tegola/releases/download/v0.9.0/tegola_linux_amd64.zip
      unzip tegola.zip

      # Create postgresql user named "user" (?)

      # Set up sample PostGIS Bonn data
      # sudo -u postgres psql -c "CREATE USER user;"
      sudo -u postgres createdb bonn
      wget -q -O bonn_osm.sql.tar.gz https://github.com/go-spatial/tegola-example-data/raw/master/bonn_osm.sql.tar.gz
      tar xzf bonn_osm.sql.tar.gz
      sudo -u postgres psql bonn < bonn_osm.dump > /dev/null 2> /dev/null

      # Set up tegola's db access
      # permissions stills setup wrong, prolly need sudo -u
      sudo -u postgres psql -c "CREATE USER tegola;"
      sudo -u postgres psql -c "ALTER USER tegola with password 'cos316';"
      sudo -u postgres psql -d bonn -c "GRANT SELECT ON ALL TABLES IN SCHEMA public TO tegola;"

      # Add login credentials for
      # echo "local all tegola md5" >> /etc/postgresql/10/main/pg_hba.conf
      sudo sed -i.bak '
        /# DO NOT DISABLE/ i\
        local   all             tegola                                 md5
      ' /etc/postgresql/10/main/pg_hba.conf
      sudo service postgresql restart
    else
      echo "Found existing tegola data. Skipping..."
    fi

    disp "All done!"

    # Start in /vagrant instead of /home/vagrant
    if ! grep -Fxq "cd /vagrant" /home/vagrant/.bashrc
    then
    echo "cd /vagrant" >> /home/vagrant/.bashrc
    fi

  SHELL

  ## Provisioning to do on each "vagrant up"
  config.vm.provision "shell", run: "always", inline: <<-SHELL
    # Add any required provisioning here:
  SHELL

  ## CPU & RAM
  config.vm.provider "virtualbox" do |vb|
    vb.customize ["modifyvm", :id, "--cpuexecutioncap", "100"]
    vb.memory = 2048
    vb.cpus = 1

    # change ubuntu-bionic-18.04-cloudimg-console.log file location
    # vb.customize [ "modifyvm", :id, "--uartmode1", "file", File.join(Dir.pwd, "logs/ubuntu-xenial-16.04-cloudimg-console.log") ]

    # disable generating ubuntu-bionic-18.04-cloudimg-console.log file in the shared folder
    vb.customize [ "modifyvm", :id, "--uartmode1", "disconnected" ]
  end

end
