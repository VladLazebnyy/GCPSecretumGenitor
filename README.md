# GCPSecretumGenitor
CLI interface for generating/importing/exporting GPG key.
the main purpose simplify the interaction with gpg/pgp

# Design

Tool should met next requirements: 
1. Command line interface
2. Generate gpg key
3. Generating secret gpg key (on demand)
4. Export GPG key
5. Export secret GPG key (on demand)
6. Import GPG key
7. Import secret GPG key
8. Get list of current GPG keys
9. Get configuration via file/cmd flags

# Libraries choice ratio
I want to have as less third party dependencies as possible that's why I am not using:
1. viper for configuration
2. zap/logrus/whatever