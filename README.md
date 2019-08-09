# pclog

A simple Progressive Calisthenics workout logging tool

## Data Model

### Plan Description

Plans are described by YAML files with the following structure:

```yaml

name: The Name of the Plan (e.g. Convict Conditioning)
exercises:
  - name: Exercise Name (e.g. Squats)
    progressions:
      - name: Shoulderstand Squats
        goals:
          - name: Beginner standard
            sets: 1
            reps: 10
          - name: Intermediate standard
            sets: 2
            reps: 25
      - name: Jackknife Squats
        goals:
          - name: ...
  - name: ...

```
