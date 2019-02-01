# The buddy app! 

half social experiment, half improv challenge, half friend on demand, the buddy app is an app whose goal is to help people always have a friend (or stranger) by their side. 

With the goal of eliminating awkward interactions, (or creating more), the buddy app uses location data to find the closest people near you to come to your aid based on a prompt you send out as a beacon for aid. 

Laid back enough to be used for situations of being alone at lunch, but useful enough to help people get out of potentially dangerous encounters, buddy is an alibi generator for the masses.

# Setup
To run the buddy webapp locally, all you need to do is run the docker-compose script, and ask me for the environment variables to put in docker/list.env. 

I'm using mkcert for the local signed cert, since firebase won't allow an untrusted cert to make keys.
