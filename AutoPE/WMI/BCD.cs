using System;
using System.Management;

namespace AutoPE.WMI
{
    class BcdConsts
    {
        // Known GUIDS
        public const string GUID_WINDOWS_BOOTMGR = "{9dea862c-5cdd-4e70-acc1-f32b344d4795}";
        public const string GUID_DEBUGGER_SETTINGS_GROUP = "{4636856e-540f-4170-a130-a84776f4c654}";
        public const string GUID_CURRENT_BOOT_ENTRY = "{fa926493-6f1c-4193-a414-58f0b2456d1e}";
        public const string GUID_WINDOWS_LEGACY_NTLDR = "{466f5a88-0af2-4f76-9038-095b170dc21c}";

        // Known Types

        public const UInt32 BCDE_DEVICE_TYPE_BOOT_DEVICE = 0x00000001u;
        public const UInt32 BCDE_DEVICE_TYPE_PARTITION = 0x00000002u;
        public const UInt32 BCDE_DEVICE_TYPE_FILE = 0x00000003u;
        public const UInt32 BCDE_DEVICE_TYPE_RAMDISK = 0x00000004u;

        public const UInt32 BCDE_LIBRARY_TYPE_APPLICATIONPATH = 0x12000002u;
        public const UInt32 BCDE_LIBRARY_TYPE_APPLICATIONDEVICE = 0x11000001u;
        public const UInt32 BCDE_LIBRARY_TYPE_DESCRIPTION = 0x12000004u;

        public const UInt32 BCDO_TYPE_VISTA_OS_ENTRY = 0x10200003u;
        public const UInt32 BCDE_OSLOADER_OSDEVICE = 0x21000001u;
        public const UInt32 BCDE_OSLOADER_SYSTEMROOT = 0x22000002u;
        public const UInt32 BCDE_OSLOADER_WINPEMODE = 0x26000022u;

        public const UInt32 BCDE_BOOTMGR_BOOTSEQUENCE = 0x24000002u;

        public const UInt32 BCDO_TYPE_DEVICE = 0x30000000u;
        public const UInt32 BCDE_DEVICEOBJ_TYPE_SDIDEVICE = 0x31000003;
        public const UInt32 BCDE_DEVICEOBJ_TYPE_SDIPATH = 0x32000004;

    }
    class BcdStore : WMIObject
    {
        public BcdStore(ManagementObject rawObject) : base(rawObject)
        {
        }

        public BcdStore(string filePath) : this(new ManagementObject($"root\\WMI:BcdStore.FilePath=\"{filePath}\""))
        {
        }

        public BcdStore() : this("")
        {
        }

        public bool CreateObject(string id, UInt32 type) =>
            CallMethod<bool>("CreateObject", new
            {
                Id = id,
                Type = type
            });
    }

    class BcdObject : WMIObject
    {
        public BcdObject(ManagementObject rawObject) : base(rawObject)
        {
        }

        public BcdObject(string id, string storeFilePath = "") : this(new ManagementObject($"root\\WMI:BcdObject.Id=\"{ id }\",StoreFilePath=\"{ storeFilePath }\""))
        {
        }

        public bool SetBooleanElement(UInt32 type, bool boolean)
            => CallMethod<bool>("SetBooleanElement", new
            {
                Type = type,
                Boolean = boolean
            });

        public bool SetStringElement(UInt32 type, string str)
            => CallMethod<bool>("SetStringElement", new
            {
                Type = type,
                String = str
            });

        public bool SetFileDeviceElement(UInt32 type, UInt32 deviceType, string additionalOpts, string path, UInt32 parentDeviceType, string parentAdditionalOpts, string parentPath)
            => CallMethod<bool>("SetFileDeviceElement", new
            {

                Type = type,
                DeviceType = deviceType,
                AdditionalOptions = additionalOpts,
                Path = path,
                ParentDeviceType = parentDeviceType,
                ParentAdditionalOptions = parentAdditionalOpts,
                ParentPath = parentPath
            });

        public bool SetPartitionDeviceElement(UInt32 type, UInt32 deviceType, string additionalOpts, string path)
            => CallMethod<bool>("SetPartitionDeviceElement", new
            {
                Type = type,
                DeviceType = deviceType,
                AdditionalOptions = additionalOpts,
                Path = path
            });

        public bool SetObjectListElement(UInt32 type, string[] ids)
            => CallMethod<bool>("SetObjectListElement", new
            {
                Type = type,
                Ids = ids
            });
    }
}
