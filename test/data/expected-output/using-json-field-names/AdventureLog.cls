/**
* DO NOT EDIT
* Generated by protoc-gen-salesforce-apex
* Source Proto File: adventure_log.proto 
* Source Message: AdventureLog
*/

global class AdventureLog {
    global enum TerrainType {
        UNKNOWN_TERRAIN,
        MOUNTAIN,
        FOREST,
        RIVER
    }
    global enum TripStatus {
        PLANNED,
        ONGOING,
        COMPLETED,
        CANCELLED,
        POSTPONED
    }
    global String tripName;
    global AdventureLog.TerrainType terrain;
    global Integer durationInHours;
    global Integer routeMap;
    global Integer participantCount;
    global Boolean isEcoFriendly;
    global Double difficultyRating;
    global List<Integer> waypointIds;
    global List<Integer> elevationChanges;
    global List<AdventureLog.TerrainType> observedTerrains;
    global List<Object> wildlifeSightings;
    global List<AdventureLog> relatedTrips;
    global List<ExplorerNote> notes;
    global Map<String, AdventureLog.CampsiteInfo> campsiteDetails;
    global Map<String, String> travelDiary;
    global AdventureLog.CampsiteInfo baseCampInfo;
    global Object specialGearList;
    global AdventureLog sisterTrip;
    global class ElevationMap {
        global List<AdventureLog.ElevationMap.ElevationByPointEntry> elevationByPoint;
    }
    global class AdventureCollection {
        global List<AdventureLog.ElevationMap> elevationProfiles;
    }
    global class CampsiteInfo {
        global String campName;
        global Boolean hasFirePit;
    }
    global class Trail {
        global String name;
        global String length;
    }
}
