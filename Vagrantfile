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
  config.vm.box = "ubuntu/trusty64"

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
    # install programs needed for setup + testing
    sudo apt-get install -y toilet
    sudo apt-get install -y jq
    disp() { toilet -f mono9 $1; }
    disp "Welcome to COS316"

    # disp "APT Updates"

    # Add MySQL repository to APT
    MYSQL_REPO=http://repo.mysql.com/apt/ubuntu/
    MYSQL_SOURCES=/etc/apt/sources.list.d/mysql.list
    touch "${MYSQL_SOURCES}"
    if grep -q "${MYSQL_REPO}" "${MYSQL_SOURCES}"; then
      echo "MySQL repo already up to date. Skipping..."
    else
      echo "Adding MySQL repo..."
      sudo apt-key add /vagrant/config_files/mysql-key
      echo "deb ${MYSQL_REPO} trusty mysql-5.7" >> "${MYSQL_SOURCES}"
    fi

    # Add PostgreSQL repo to APT
    POSTGRES_REPO=http://apt.postgresql.org/pub/repos/apt
    POSTGRES_SOURCES=/etc/apt/sources.list
    if grep -q "${POSTGRES_REPO}" "${POSTGRES_SOURCES}"; then
      echo "PostgreSQL repo already up to date. Skipping..."
    else
      echo "Adding PostgreSQL repo..."
      echo "deb ${POSTGRES_REPO} trusty-pgdg main" >> ${POSTGRES_SOURCES}
      wget --quiet -O - http://apt.postgresql.org/pub/repos/apt/ACCC4CF8.asc | sudo apt-key add -
    fi

    # COS 316 Assignments
    # Assignment 1
    sudo add-apt-repository ppa:gophers/archive # For Go 1.11
    sudo apt-get update
    sudo apt-get -y upgrade
    disp "Emacs"
    sudo apt-get install -y emacs


    # echo "export PYTHONPATH=${PYTHONPATH}:/vagrant/course-bin" >> /home/vagrant/.profile

    # Set correct permissions for bash scripts
    find /vagrant -name "*.sh" | xargs chmod -v 744

    # If the repository was pulled from Windows, convert line breaks to Unix-style
    sudo apt-get install -y dos2unix
    printf "Using dos2unix to convert files to Unix format if necessary..."
    find /vagrant -name "*" -type f | xargs dos2unix -q

    # Install Go 1.11, the most recent for Ubuntu 14.04
    disp "Go v1.11"
    sudo apt-get install -y golang-1.11-go
    echo "export PATH=${PATH}:/usr/lib/go-1.11/bin:/go/bin" >> /home/vagrant/.profile
    echo "export GOPATH='/go'" >> /home/vagrant/.profile
    # echo "export GO111MODULE=auto" >> /home/vagrant/.profile
    source /home/vagrant/.profile # Make sure go environment variables are set

    # Install MySQL (for Tinode), skipping interactive configuration
    # Note: The default password is obviously not secure. In a real application
    # we would want to change it, but for our purposes this is "good enough"
    disp "MySQL 5.7"
    MYSQL_PWD="cos316"
    sudo debconf-set-selections <<< "mysql-server mysql-server/root_password password ${MYSQL_PWD}"
    sudo debconf-set-selections <<< "mysql-server mysql-server/root_password_again password ${MYSQL_PWD}"
    sudo debconf-set-selections <<< "mysql-community-server mysql-community-server/root-pass password ${MYSQL_PWD}"
    sudo debconf-set-selections <<< "mysql-community-server mysql-community-server/re-root-pass password ${MYSQL_PWD}"
    sudo apt-get -y install mysql-server

    # Download and install Tinode chat server
    disp "Tinode"
    sudo apt-get install -y git
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
    sudo apt-get install -y unzip
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

    # Install PostgreSQL and PostGIS
    disp "PostgreSQL"
    sudo apt-get install -y postgresql-10
    sudo apt-get install -y postgresql-10-postgis-2.4
    sudo apt-get install -y postgresql-10-postgis-2.4-scripts

    # Install Tegola
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
      sudo -u postgres psql bonn < bonn_osm.dump

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

    # COS 461 Assignments
    # Assignment 1
    #sudo apt-get install -y python-dev
    #curl -sS https://bootstrap.pypa.io/get-pip.py > get-pip.py
    #sudo python get-pip.py
    #rm -f get-pip.py
    # Install old version of tornado before installing jupyter
    #sudo pip install tornado==4.5.3
    #sudo pip install jupyter
    #sudo pip install -U tzupdate

    # Assignment 2
    #sudo apt-get install -y python-tk

    # Assignment 3
    #sudo pip install nbconvert
    #sudo apt-get install -y mininet
    #sudo apt-get install -y python-numpy
    #sudo apt-get install -y python-matplotlib

    # Assignment 4
    #sudo apt-get install -y whois
    #sudo pip install ipaddress

    # Assignment 6
    #sudo pip install scapy
    #sudo apt-get install -y python-scipy
    #sudo apt-get install -y bind9
    #sudo cp /vagrant/assignment6/bind/* /etc/bind

    # Assignment 7
    # sudo apt-get install -y apache2-utils
    # echo "export GOPATH=/vagrant/assignment1" >> /home/vagrant/.profile

    disp "All done!"

    # Start in /vagrant instead of /home/vagrant
    if ! grep -Fxq "cd /vagrant" /home/vagrant/.bashrc
    then
    echo "cd /vagrant" >> /home/vagrant/.bashrc
    fi
  SHELL

  ## Provisioning to do on each "vagrant up"
  config.vm.provision "shell", run: "always", inline: <<-SHELL
    sudo tzupdate 2> /dev/null
    # Assignment 3
    sudo modprobe tcp_probe port=5001 full=1
  SHELL

  ## CPU & RAM
  config.vm.provider "virtualbox" do |vb|
    vb.customize ["modifyvm", :id, "--cpuexecutioncap", "100"]
    vb.memory = 2048
    vb.cpus = 1
  end

end
