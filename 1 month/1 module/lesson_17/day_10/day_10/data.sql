-- Patient Table
CREATE TABLE Patient (
    PatientID INT PRIMARY KEY,
    Name VARCHAR(50),
    Surname VARCHAR(50),
    Phone VARCHAR(15)
);

-- Cashier Table
CREATE TABLE Cashier (
    CashierID INT PRIMARY KEY,
    Name VARCHAR(50),
    Surname VARCHAR(50),
    Phone VARCHAR(15),
    ClinicID INT,
    FOREIGN KEY (ClinicID) REFERENCES Clinic(ClinicID)
);

-- Doctor Table
CREATE TABLE Doctor (
    DoctorID INT PRIMARY KEY,
    Name VARCHAR(50),
    Surname VARCHAR(50),
    Phone VARCHAR(15),
    ClinicID INT,
    SpecializationID INT,
    FOREIGN KEY (ClinicID) REFERENCES Clinic(ClinicID),
    FOREIGN KEY (SpecializationID) REFERENCES Specialization(SpecializationID)
);

-- Clinic Table
CREATE TABLE Clinic (
    ClinicID INT PRIMARY KEY,
    Name VARCHAR(100),
    Logo VARCHAR(255),
    Ranking INT,
    City VARCHAR(50),
    BranchOfficeQuantity INT
);

-- Specialization Table
CREATE TABLE Specialization (
    SpecializationID INT PRIMARY KEY,
    SpecializationName VARCHAR(50)
);

-- BranchOffice Table
CREATE TABLE BranchOffice (
    BranchOfficeID INT PRIMARY KEY,
    ClinicID INT,
    Address VARCHAR(255),
    Phone VARCHAR(15),
    FOREIGN KEY (ClinicID) REFERENCES Clinic(ClinicID)
);