db = db.getSiblingDB('rust_drop');
db.createCollection("run_time")

db.run_time.insert({
    started_at: new Date()
})